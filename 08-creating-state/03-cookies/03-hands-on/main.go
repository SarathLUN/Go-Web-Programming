package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("my-cookie")
	if err != nil {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "1",
			Path:  "/",
		}
	} else {
		count, err := strconv.Atoi(cookie.Value) // ASCI to Int
		if err != nil {
			log.Fatalln(err)
		}
		count++
		cookie.Value = strconv.Itoa(count) // Int to ASCI
	}

	http.SetCookie(w, cookie)
	_, err = io.WriteString(w, cookie.Value)
	if err != nil {
		log.Println(err)
	}
}
