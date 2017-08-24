package view

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tpIndex = template.New("")
)

func init() {
	tpIndex.Funcs(template.FuncMap{})
	_, err := tpIndex.ParseFiles("template/root.tmpl", "template/index.tmpl")

	if err != nil {
		panic(err)
	}
	tpIndex = tpIndex.Lookup("root")
}

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := tpIndex.ParseFiles(files...)

	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; chartset=utf-8")
	err := t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

// Index Renders index View
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}
