package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/test2")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Println("error close db connection:", err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func drop(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	stm, err := db.Prepare(`DROP TABLE customers;`)
	if err != nil {
		log.Println(err)
	}
	defer stm.Close()

	r, err := stm.Exec()
	if err != nil {
		log.Println(err)
	}

	n, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, "customers dropped:", n)
}

func del(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	stm, err := db.Prepare(`DELETE FROM customers WHERE name='new name';`)
	if err != nil {
		log.Println(err)
	}
	defer stm.Close()

	r, err := stm.Exec()
	if err != nil {
		log.Println(err)
	}

	n, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, "record deleted:", n)
}

func update(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	stm, err := db.Prepare(`UPDATE customers SET name='new name' WHERE name='Tony';`)
	if err != nil {
		log.Println(err)
	}
	defer stm.Close()

	r, err := stm.Exec()
	if err != nil {
		log.Println(err)
	}

	n, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, "updated record:", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	rows, err := db.Query(`SELECT * FROM customers;`)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var name, output string

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
		}
		output += name + ", "
	}
	fmt.Fprintln(w, output)
}

func insert(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	stm, err := db.Prepare(`INSERT INTO customers value ('Tony')`)
	if err != nil {
		log.Println(err)
	}
	defer stm.Close()

	r, err := stm.Exec()
	if err != nil {
		log.Println(err)
	}

	n, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, "added Tony:", n)
}

func create(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	stm, err := db.Prepare(`CREATE TABLE customers (name VARCHAR(20));`)
	if err != nil {
		log.Fatalln(err)
	}
	defer stm.Close()

	result, err := stm.Exec()
	if err != nil {
		log.Fatalln(err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(w, "TABLE customers CREATED:", n)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	rows, err := db.Query(`SELECT aName FROM amigos;`)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(rows)

	var s, name string
	s = "RETRIEVED RECORDS:\n"
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Fatalln(err)
		}
		s += name + "\n"
	}
	_, err = fmt.Fprintln(w, s)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	io.WriteString(w, `
		<a href="/amigos">list out amigos</a> <br>
		<a href="/create">create table customers</a> <br>
		<a href="/insert">insert into customers</a> <br>
		<a href="/read">read from customers</a> <br>
		<a href="/update">update customer</a> <br>
		<a href="/delete">delete customer</a> <br>
		<a href="/drop">drop table customers</a> <br>
`)
}
