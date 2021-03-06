package main

import (
	"encoding/json"
	"log"
)

// We can choose to only unmarshal some of the json data
// create a data struct that only has fields we need
type city struct {
	Latitude, Longitude float64
	City                string
}

type cities []city

func main() {
	var data cities
	rcvd := `[{"precision":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"precision":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
}
