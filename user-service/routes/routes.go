package routes

import (
	"go-microservices-app/user-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/users", controllers.CreateUser).Methods("POST")
}
