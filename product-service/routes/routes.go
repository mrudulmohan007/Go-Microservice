package routes

import (
	"go-microservices-app/product-service/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/", controllers.ServeHome).Methods("GET")
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.GetProductById).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/product/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", controllers.DeleteProduct).Methods("DELETE")
}
