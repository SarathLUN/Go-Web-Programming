package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := region{
		Region: "Southern",
		Hotels: []hotel{
			{
				"Hotel California",
				"42 Sunset Boulevard",
				"Los Angeles",
				"95612",
			},
			{
				"H",
				"4",
				"L",
				"8765",
			},
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
