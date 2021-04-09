package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "from hotdog")
}

type hotcat int

func (h hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "from hotcat")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog", d)
	http.Handle("/cat", c)

	// apply default serve mux by nil handler
	http.ListenAndServe(":8080", nil)
}
