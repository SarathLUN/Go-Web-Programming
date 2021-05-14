package main

import (
	"encoding/json"
	"log"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{}
	log.Println(m)
	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(bs))
}
