package controllers

import (
	"encoding/json"
	"go-microservices-app/product-service/models"
	"net/http"

	"github.com/gorilla/mux"
)

var products []models.Product

func init() {
	// Seed with some dummy products
	products = append(products, models.Product{ID: "1", Name: "Laptop", Price: 1000, UserID: "1"})
	products = append(products, models.Product{ID: "2", Name: "Phone", Price: 500, UserID: "2"})
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, product := range products {
		if product.ID == params["id"] {
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	// Add the product to the slice
	products = append(products, product)

	// Return the newly created product
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}
