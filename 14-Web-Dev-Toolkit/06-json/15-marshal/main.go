package main

import (
	"encoding/json"
	"log"
	"os"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{
		true,
		[]string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}
	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatalln(err)
	}
	os.Stdout.Write(bs)
}

// Answer
// "... Struct values encode as JSON objects.
// Each exported struct field becomes a member of the object..."
