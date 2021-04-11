package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
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
	var s string
	if r.Method == http.MethodPost {

		// open file
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(f multipart.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)

		// FYI
		fmt.Println("file:", f)
		fmt.Println("header:", h)

		// read file
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)

		// store file
		dst, err := os.Create(filepath.Join("./uploads/", h.Filename))
		if err != nil {
			log.Fatalln(err)
		}
		defer func(dst *os.File) {
			err := dst.Close()
			if err != nil {
				log.Fatalln(err)
			}
		}(dst)
		_, err = dst.Write(bs)
		if err != nil {
			log.Fatalln(err)
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	if err != nil {
		log.Fatalln(err)
	}
}
