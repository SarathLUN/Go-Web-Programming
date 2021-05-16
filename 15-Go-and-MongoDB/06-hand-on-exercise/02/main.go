package main

import (
	"github.com/SarathLUN/Go-Web-Programming/15-Go-and-MongoDB/06-hand-on-exercise/02/controllers"
	"github.com/SarathLUN/Go-Web-Programming/15-Go-and-MongoDB/06-hand-on-exercise/02/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}
