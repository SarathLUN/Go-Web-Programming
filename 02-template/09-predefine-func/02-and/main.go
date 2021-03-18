package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	users := []user{
		{
			"Buddha",
			"The belief of no beliefs",
			false,
		},
		{
			"Gandhi",
			"Be the change",
			true,
		},
		{
			"",
			"Nobody",
			true,
		},
	}
	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
