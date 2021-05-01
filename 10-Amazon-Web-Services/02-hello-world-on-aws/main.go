package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Method)
	io.WriteString(w, "Oh yeah, I'm running on AWS.")
}