package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Tony-Key", "Hello World!")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>This write from HotDog</h1>")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
