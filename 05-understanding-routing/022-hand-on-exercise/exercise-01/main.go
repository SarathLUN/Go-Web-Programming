package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", bar)
	http.HandleFunc("/me", myName)
	http.ListenAndServe(":8080", nil)
}

func myName(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, Tony!")
}

func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "bar ran")
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}
