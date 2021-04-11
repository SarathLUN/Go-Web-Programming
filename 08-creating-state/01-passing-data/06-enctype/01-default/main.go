package main

import (
	"html/template"
	"log"
	"net/http"
)

// parse the template
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
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	// body
	bs := make([]byte, r.ContentLength)
	_, err := r.Body.Read(bs)
	if err != nil {
		log.Fatalln(err)
	}
	body := string(bs)

	err = tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		log.Fatalln(err)
	}
}
