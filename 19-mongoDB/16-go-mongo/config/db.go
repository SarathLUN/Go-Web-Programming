package config

import (
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
	"log"
)

var DB *mgo.Database

var Books *mgo.Collection

func init() {
	// get a mongo sessions
	s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost:27017/bookstore")
	if err != nil {
		panic(err)
	}
	if err = s.Ping(); err != nil {
		panic(err)
	}
	DB = s.DB("bookstore")
	Books = DB.C("books")
	log.Println("You connected to your mongo database.")
}
