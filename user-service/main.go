package main

import (
	"go-microservices-app/user-service/models"
	"go-microservices-app/user-service/routes"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var users []models.User // Declare users variable in the main package

func main() {

	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	log.Println("User Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
