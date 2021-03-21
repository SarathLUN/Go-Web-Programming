package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Description, Meal string
	Price                   float64
}

type items []item

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := items{
		item{
			Name:        "Oatmeal",
			Description: "yum yum",
			Meal:        "Breakfast",
			Price:       4.95,
		},
		item{
			Name:        "Hamburger",
			Description: "Delicous good eating for you",
			Meal:        "Lunch",
			Price:       6.95,
		},
		item{
			Name:        "Pasta Bolognese",
			Description: "From Italy delicious eating",
			Meal:        "Dinner",
			Price:       7.95,
		},
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
