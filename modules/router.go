package modules

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"breakfast/config"

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

	addr := fmt.Sprintf("%s:%d", config.Main.WebServer.Interface, config.Main.WebServer.Port)

	log.Println("listening on "+addr+"...")

	panic(http.ListenAndServe(addr, r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if r.URL.Query().Has("short") && shortTemplate != nil {
		log.Println("handling short-response request: "+r.URL.Path)
		shortTemplate.Execute(w, templateData)
	} else {
		log.Println("handling request: "+r.URL.Path)
		longTemplate.Execute(w, templateData)
	}
}
