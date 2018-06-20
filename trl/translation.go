// Package trl is a central store for application environment translations;
// While specific objects such as QuestionaireT contain translations for its contents,
// we need some global store too.
package trl

// LangCodes for returning multiple translations.
// When no langCode is available, then the first entry rules.
// A call to All() returns explicitly all key-values.
// LangCodes will be initialized in cfg.Load().LangCodes; we prevent circular dependency
var LangCodes = []string{"de", "en"}

// S stores a multi lingual string.
// Contains one value for each language code.
type S map[string]string

const noTrans = "multi lingual string not initialized."

// Tr translates by key.
// Defaults           to english.
// Defaults otherwise to first existing map key.
// Returns a 'signifiers string' if no translation exists,
// to help to uncover missing translations.
func (s S) Tr(langCode string) string {
	if val, ok := s[langCode]; ok {
		return val
	}
	if val, ok := s[LangCodes[0]]; ok {
		return val
	}
	if s == nil {
		return noTrans
	}
	return noTrans
}

// TrSilent gives no warning - if the translation is not set.
// Good if we do not require a translation string.
// Good for i.e. HTML title attribute - where errors are easy to overlook.
func (s S) TrSilent(langCode string) string {
	ret := s.Tr(langCode)
	if ret == noTrans {
		return ""
	}
	return ret
}

// String is the default "stringer" implementation
func (s S) String() string {
	if val, ok := s["en"]; ok {
		return val
	}
	if val, ok := s[LangCodes[0]]; ok {
		return val
	}
	for _, val := range s {
		return val
	}
	return ""
}

// All returns all translations
func (s S) All() string {
	ret := ""
	for _, key := range LangCodes {
		if val, ok := s[key]; ok {
			ret += val
			ret += "\n\n"
		}
	}
	return ret
}

// Map - Translations Type
// Usage in templates
// 		{{.Trls.imprint.en                     }}  // directly accessin a specific translation; chaining the map keys
// 		{{.Trls.imprint.Tr       .Sess.LangCode}}  // using .Tr(langCode)
// 		{{.Trls.imprint.TrSilent .Sess.LangCode}}  //
type Map map[string]S