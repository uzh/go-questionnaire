// Package cfg implements a configuration database,
// loaded from a json file.
// Filename must be given as command line argument
// or environment variable.
// Access to the config data is made in threadsafe manner.
package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/zew/util"
)

type ConfigT struct {
	sync.Mutex

	IsProduction bool `json:"is_production"` // true => templates are not recompiled

	AppName       string `json:"app_name"`       // with case, i.e. Taxkit
	UrlPathPrefix string `json:"urlpath_prefix"` // lower case - no slashes, i.e. taxkit
	AppMnemonic   string `json:"app_mnemonic"`   // For differentiation of static dirs - when UrlPathPrefix is empty; imagine multiple instances

	BindHost               string `json:"bind_host"`
	BindSocket             string `json:"bind_socket"`
	BindSocketFallbackHttp string `json:"bind_socket_fallback_http"`
	Tls                    bool   `json:"tls"`
	Tls13                  bool   `json:"tls13"`               // ultra safe - but excludes internet explorer 11
	HttpReadTimeOut        int    `json:"http_read_time_out"`  // for large requests
	HttpWriteTimeOut       int    `json:"http_write_time_out"` // for *sending* large files over slow networks, i.e. ula's videos, set to 30 or 60 secs

	Css map[string]string `json:"css"` // differentiate multiple instances by color and stuff - without duplicating entire css files

}

// Obtained by ENV variable or command line flag in main package.
// Being set from the main package.
// Holds the relative path and filename to look for; could be ".cfg/config.json".
// Relative to the app main dir.
var CfgPath = path.Join(".", "config.json")

var cfgS *ConfigT // package variable 'singleton' - needs to be an allocated struct - to hold pointer receiver-re-assignment

func Get() *ConfigT {
	// Same as lgn.Get().
	// No lock needed here.
	// Since in load(), we simply exchange one pointer by another at the end of loading.
	// c.Lock()
	// defer c.Unlock()
	return cfgS
}

// No method to ConfigT, no pointer receiver;
// We could only *copy*:  *c = *newCfg
func Load() {
	// c.Lock()
	// defer c.Unlock()
	file, err := util.LoadConfigFile(CfgPath)
	if err != nil {
		log.Fatalf("Could not load config file: %v", err)
	}
	log.Printf("Found config file: %v", CfgPath)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Err closing config file: %v", err)
		}
		log.Printf("Closed config file: %v", CfgPath)
	}()

	decoder := json.NewDecoder(file)
	tempCfg := ConfigT{}
	err = decoder.Decode(&tempCfg)
	if err != nil {
		log.Fatal(err)
	}

	if tempCfg.AppName == "" {
		log.Fatal("Config underspecified; at least app_name should be set")
	}
	if tempCfg.AppMnemonic == "" {
		log.Fatal("Config underspecified; at least app_mnemonic should be set")
	}
	cfgS = &tempCfg // replace pointer in one go - should be threadsafe
	log.Printf("config loaded 1\n%#s", util.IndentedDump(cfgS))
}

//
func LoadH(w http.ResponseWriter, r *http.Request) {
	Load()
	w.Write([]byte("cfg reloaded"))
}

func (c *ConfigT) Save(fn ...string) {

	firstColLeftMostPrefix := " "
	byts, err := json.MarshalIndent(c, firstColLeftMostPrefix, "\t")
	util.BubbleUp(err)

	saveDir := path.Dir(CfgPath)
	err = os.Chmod(saveDir, 0755)
	util.BubbleUp(err)

	configFile := path.Base(CfgPath)
	if len(fn) > 0 {
		configFile = fn[0]
	}
	savePath := path.Join(saveDir, configFile)
	err = ioutil.WriteFile(savePath, byts, 0644)
	util.BubbleUp(err)

	log.Printf("Saved config file to %v", savePath)
}

// Any URL Path is prefixed with the UrlPathPrefix, if UrlPathPrefix is set.
// Prevents unnecessary slashes.
// No trailing slash
// Routes with trailing "/" such as "/path/"
// get a redirect "/path" => "/path/" if "/path" is not registered yet.
// This behavior of func server.go - (mux *ServeMux) Handle(...) is nasty
// since it depends on the ORDER of registrations.
//
// Best strategy might be
//    mux.HandleFunc(appcfg.Pref(urlPath), argFunc)      // Claim "/path"
//    mux.HandleFunc(appcfg.PrefWTS(urlPath), argFunc)   // Claim "/path/"
// Notice the order - other way around would block "/path" with a redirect handler
func Pref(pth ...string) string {

	if cfgS.UrlPathPrefix != "" {
		if len(pth) > 0 {
			return path.Join("/", cfgS.UrlPathPrefix, pth[0])
		}
		return path.Join("/", cfgS.UrlPathPrefix)
	}

	// No UrlPathPrefix
	if len(pth) > 0 {
		return path.Join("/", pth[0])
	}
	return ""

}

// Prefix with trailing slash
func PrefWTS(pth ...string) string {
	p := Pref(pth...)
	return p + "/"
}

func init() {
	ex := &ConfigT{
		IsProduction:           false,
		AppName:                "My Example App Label",
		UrlPathPrefix:          "exmpl",
		AppMnemonic:            "exmpl",
		BindHost:               "0.0.0.0",
		BindSocket:             "8081",
		BindSocketFallbackHttp: "8082",
		Tls:              false,
		Tls13:            false,
		HttpReadTimeOut:  5,
		HttpWriteTimeOut: 30,
	}
	ex.Save("config-example.json")
}
