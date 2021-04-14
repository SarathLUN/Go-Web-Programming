package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred:", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar:", r.Method)
	w.Header().Set("Location", "/barred")
	w.WriteHeader(http.StatusSeeOther)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo:", r.Method)
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
