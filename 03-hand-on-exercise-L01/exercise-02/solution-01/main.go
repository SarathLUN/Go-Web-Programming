package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip, Region string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			"Hotel California",
			"42 Sunset Boulevard",
			"Los Angeles",
			"95612",
			"southern",
		},
		hotel{
			"H",
			"4",
			"L",
			"847774",
			"southern",
		},
	}
	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
