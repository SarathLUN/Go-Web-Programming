package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method,
		req.Form,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
