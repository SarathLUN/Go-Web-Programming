package main

import (
	"github.com/SarathLUN/Go-Web-Programming/15-Go-and-MongoDB/05-CRUD-with-Go-Mongodb/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func init() {

}

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}
	return s
}
