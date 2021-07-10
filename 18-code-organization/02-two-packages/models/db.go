package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	log.Println("You connected to your database.")
}
