package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})
	_, err := fmt.Fprintln(w, "Cookies written - check your browser")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
	if err != nil {
		log.Fatalln(err)
	}
}

func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		log.Println("Cookie #1:", err)
	} else {
		_, err := fmt.Fprintln(w, "Your Cookie #1:", c1)
		if err != nil {
			log.Fatalln(err)
		}
	}
	c2, err := r.Cookie("general")
	if err != nil {
		log.Println("Cookie #2:", err)
	} else {
		_, err := fmt.Fprintln(w, "Your Cookie #2:", c2)
		if err != nil {
			log.Fatalln(err)
		}
	}
	c3, err := r.Cookie("specific")
	if err != nil {
		log.Println("Cookie #3:", err)
	} else {
		_, err := fmt.Fprintln(w, "Your Cookie #3:", c3)
		if err != nil {
			log.Fatalln(err)
		}
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
