package main

import (
	"encoding/json"
	"log"
)

func main() {
	var data int
	rcvd := `66`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
	log.Printf("%T", data)
}
