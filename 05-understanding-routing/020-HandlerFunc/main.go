package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog")
}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat")
}

func main() {
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
