package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar:", r.Method)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at foo:", r.Method)
	http.Redirect(w, r, "/bar", http.StatusMovedPermanently)
}
