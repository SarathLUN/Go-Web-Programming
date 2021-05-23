package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type Books struct {
	isbn   string
	title  string
	author string
	price  float32
}

func main() {
	db, err := sql.Open("postgres", "postgresql://bond:password@localhost:5432/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	bks := make([]Books, 0)
	for rows.Next() {
		bk := Books{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price) // make sure in correct order
		if err != nil {
			log.Println(err)
		}
		bks = append(bks, bk)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

	log.Println(bks)
}
