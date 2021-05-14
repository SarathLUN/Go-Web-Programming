package main

import (
	"encoding/json"
	"log"
)

func main() {
	var a []string
	rcvd := `null`
	err := json.Unmarshal([]byte(rcvd), &a)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(a)
	log.Println(len(a))
	log.Println(cap(a))
}
