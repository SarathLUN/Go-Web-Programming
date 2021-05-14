package main

import (
	"encoding/json"
	"log"
)

type thumbnail struct {
	URL           string
	Height, Width int
}

type img struct {
	Width, Height int
	Title         string
	Thumbnail     thumbnail
	Animated      bool
	IDs           []int
}

func main() {
	var data img
	rcvd := `
	{
		"Width":800,
		"Height":600,
		"Title":"View from 15th Floor",
		"Thumbnail":{
			"Url":"http://www.example.com/image/481989943",
			"Height":125,
			"Width":100
			},
		"Animated":false,
		"IDs":[116,943,234,38793]
}
`
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(data)
	for i, v := range data.IDs {
		log.Println("id#", i, "=", v)
	}
	log.Println(data.Thumbnail.URL)
}
