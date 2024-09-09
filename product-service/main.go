package main

import (
	"go-microservices-app/product-service/routes"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterProductRoutes(r)
	log.Println("Product Service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
