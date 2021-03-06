package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName    string
	LastName     string
	IsSubscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)
	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		log.Fatalln(err)
	}
}
