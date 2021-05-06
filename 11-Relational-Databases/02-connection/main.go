package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	// for AWS
	// db, err = sql.Open("mysql", "admin:yzTvwnqe6AJ6CYb@tcp(database-1.ch3fcoesrgva.us-east-2.rds.amazonaws.com:3306)/test2?charset=utf8")
	// my localhost
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test2?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)

}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "Successful")
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
	log.Println("no error, mean connected db")
}
