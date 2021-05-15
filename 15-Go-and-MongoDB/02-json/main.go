package main

import (
	"encoding/json"
	"fmt"
	"github.com/SarathLUN/Go-Web-Programming/15-Go-and-MongoDB/02-json/models"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	// route with param
	r.GET("/user/:id", getUser)
	http.ListenAndServe(":8080", r)
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		"James Bond",
		"Male",
		32,
		p.ByName("id"),
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}
	// response back to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
	/*output: {"name":"James Bond","gender":"Male","age":32,"id":"9872309847"}*/

	//fmt.Fprintln(w, uj)
	/*output: [123 34 110 97 109 101 34 58 34 74 97 109 101 115 32 66 111 110 100 34 44 34 103 101 110 100 101 114 34 58 34 77 97 108 101 34 44 34 97 103 101 34 58 51 50 44 34 105 100 34 58 34 57 56 55 50 51 48 57 56 52 55 34 125]*/
}
func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}
