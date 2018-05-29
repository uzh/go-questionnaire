package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zew/go-questionaire/qst"

	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/go-questionaire/tpl"
)

func errorH(w http.ResponseWriter, r *http.Request, msg string) {

	sess := sessx.New(w, r)

	log.Print(msg)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tplBundle := tpl.Get(w, r, "main.html")

	ts := &tpl.StackT{"error.html", "non-existent.html"}

	d := tplDataExtT{
		Q: &qst.QuestionaireT{LangCode: "en"}, // just setting the lang code for the outer layout template
	}
	d.TplDataT = tpl.TplDataT{
		TplBundle: tplBundle,
		TS:        ts,
		Sess:      &sess,
		Cnt:       msg,
	}

	err := tplBundle.Execute(w, d)
	if err != nil {
		s := fmt.Sprintf("Executing template caused: %v", err)
		log.Print(s)
		w.Write([]byte(s))
		return
	}
}