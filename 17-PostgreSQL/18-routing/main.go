package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Books struct {
	isbn   string
	title  string
	author string
	price  float32
}

var db *sql.DB
var err error

func init() {

	db, err = sql.Open("postgres", "postgresql://bond:password@localhost:5432/bookstore?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("You connected to your database.")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World;")
	})
	http.HandleFunc("/books", books)
	http.ListenAndServe("localhost:8080", nil)

}

func books(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	//defer db.Close()
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	bks := make([]Books, 0)
	for rows.Next() {
		bk := Books{}
		err = rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // make sure in correct order
		if err != nil {
			log.Fatalln(err)
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	_, err = fmt.Fprintln(w, bks)
	if err != nil {
		log.Fatalln(err)
	}
}
