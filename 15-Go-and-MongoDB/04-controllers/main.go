package main

import (
	"github.com/SarathLUN/Go-Web-Programming/15-Go-and-MongoDB/04-controllers/controllers"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	r := httprouter.New()
	userController := controllers.NewUserController()

	r.GET("/user/:id", userController.GetUser)
	// added route
	r.POST("/user", userController.CreateUser)
	// added route plus parameter
	r.DELETE("/user/:id", userController.DeleteUser)
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Serving: http://localhost:8080 ")
}
