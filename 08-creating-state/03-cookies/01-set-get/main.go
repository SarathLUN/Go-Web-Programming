package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	_, err = fmt.Fprintln(w, "YOUR COOKIE:", c)
	if err != nil {
		log.Fatalln(err)
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	_, err := fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
	if err != nil {
		log.Fatalln(err)
	}
}
