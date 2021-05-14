package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{
			"Suit",
			"Gun",
			"Wry sense of humor",
		},
	}
	// encode take go struct and turn into json then immediate write to response
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		"James",
		"Bond",
		[]string{
			"Suit",
			"Gun",
			"Wry sense of humor",
		},
	}
	// Marshal just return json
	// here we can handle error in step-by-step
	// we can also use return json in multiple purpose, if needed
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	// here we write to response
	_, err = w.Write(j)
	if err != nil {
		log.Println(err)
	}
}

func foo(w http.ResponseWriter, req *http.Request) {
	log.Println("You are at foo()")
	w.Write([]byte("You are at foo."))
}
