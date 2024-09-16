package controllers

import (
	"encoding/json"
	"fmt"
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

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello this is the product page guys!</h1>"))
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
	fmt.Println("Creating one product")
	w.Header().Set("Content-Type", "application/json")

	//what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	// Loop through courses to check for duplicate product
	for _, existingProduct := range products {
		if existingProduct.Name == product.Name {
			json.NewEncoder(w).Encode("Product name already exists")
			return
		}
	}

	//append book into products

	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// updating the products
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedProduct models.Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)

	for index, product := range products {
		if product.ID == params["id"] {
			// Update the product in the slice
			products[index] = updatedProduct

			// Return the updated product
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, product := range products {
		if product.ID == params["id"] {
			// Remove the product from the slice
			products = append(products[:index], products[index+1:]...)

			// Return a success message
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
