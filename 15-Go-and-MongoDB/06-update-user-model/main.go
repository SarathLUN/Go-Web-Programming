package main

import (
	"github.com/SarathLUN/Go-Web-Programming/15-Go-and-MongoDB/05-update-user-controller/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	// get a UserController instance
	userController := controllers.NewUserController(getSession())
	r.GET("/user/:id", userController.GetUser)
	r.POST("/user", userController.CreateUser)
	r.DELETE("/user/:id", userController.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost:27017/")
	if err != nil {
		log.Fatalln("mongodb error:", err)
	}
	return s
}
