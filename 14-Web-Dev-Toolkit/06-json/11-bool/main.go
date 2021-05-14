package main

import (
	"encoding/json"
	"log"
)

func main() {
	var data bool
	rcvd := `true`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
	log.Printf("%T", data)
}
