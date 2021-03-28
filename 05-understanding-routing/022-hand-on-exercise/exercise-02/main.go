package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bar)
	http.HandleFunc("/me", tony)
	http.ListenAndServe(":8080", nil)
}

func tony(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("Error parsing template:", err)
	}
	err = tpl.ExecuteTemplate(w, "something.gohtml", "Tony")
	if err != nil {
		log.Fatalln("Error executing template:", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bar ran")
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}
