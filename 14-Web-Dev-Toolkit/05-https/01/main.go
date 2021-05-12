package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("open https://localhost:10433")
	}

}

func foo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte("This is an example server.\n"))
	if err != nil {
		log.Fatalln(err)
	}
}
