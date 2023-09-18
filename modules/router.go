package modules

import (
	"fmt"
	"net/http"
	"text/template"
	"version-forge/config"

	"github.com/gorilla/mux"
)

var templateData TemplateData
var longTemplate *template.Template
var shortTemplate *template.Template

func MustListen(td TemplateData, bt *template.Template, ct *template.Template) {
	templateData = td
	longTemplate = bt
	shortTemplate = ct

	r := mux.NewRouter()
	r.HandleFunc(config.Main.WebServer.AppPath, handler)
	r.PathPrefix(config.Main.WebServer.StaticPath).Handler(http.FileServer(http.Dir(config.Main.Paths.StaticDir)))
	panic(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d", config.Main.WebServer.Interface, config.Main.WebServer.Port),
			r,
		),
	)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if r.URL.Query().Has("short") && shortTemplate != nil {
		shortTemplate.Execute(w, templateData)
	} else {
		longTemplate.Execute(w, templateData)
	}
}
