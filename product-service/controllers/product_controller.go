package controllers

import (
	"encoding/json"
	"fmt"
	"go-microservices-app/product-service/models"
	"io"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var products []models.Product

func init() {
	// Seed with some dummy products
	products = append(products, models.Product{ProductID: "1", ProductName: "Laptop", ProductPrice: 1000, UserID: "1"})
	products = append(products, models.Product{ProductID: "2", ProductName: "Phone", ProductPrice: 500, UserID: "2"})
}

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello this is the product page guys!</h1>"))
}

// Get all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// Get product by ID and fetch user details from user-service
func GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, product := range products {
		if product.ProductID == params["id"] {
			// Fetch user details from user-service based on UserID
			userDetails, err := fetchUserDetails(product.UserID)
			if err != nil {
				http.Error(w, "Failed to fetch user details", http.StatusInternalServerError)
				return
			}

			// Combine product and user details
			response := map[string]interface{}{
				"product": product,
				"user":    userDetails,
			}

			// Return combined product and user details
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Create new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one product")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	// Check for duplicate product
	for _, existingProduct := range products {
		if existingProduct.ProductName == product.ProductName {
			json.NewEncoder(w).Encode("Product name already exists")
			return
		}
	}

	// Add the new product
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

// Fetch user details from user-service
func fetchUserDetails(userID string) (map[string]interface{}, error) {
	userServiceURL := fmt.Sprintf("http://localhost:8081/users/%s", userID) //if u use docker-compose give user-service instead of this localhost
	resp, err := http.Get(userServiceURL)
	if err != nil {
		log.Printf("Error fetching user details: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user details, status code: %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	var userDetails map[string]interface{}
	json.Unmarshal(body, &userDetails)

	return userDetails, nil
}

// Update product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedProduct models.Product
	_ = json.NewDecoder(r.Body).Decode(&updatedProduct)

	for index, product := range products {
		if product.ProductID == params["id"] {
			products[index] = updatedProduct
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Delete product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, product := range products {
		if product.ProductID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
