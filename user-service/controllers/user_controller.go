package controllers

import (
	"encoding/json"
	"go-microservices-app/user-service/models"
	"net/http"

	"github.com/gorilla/mux"
)

var users []models.User // Slice to hold user data

func init() {
	// Seed with some dummy users
	users = append(users, models.User{ID: "1", Name: "John Doe", Email: "john@example.com"})
	users = append(users, models.User{ID: "2", Name: "Jane Smith", Email: "jane@example.com"})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Add the user to the slice
	users = append(users, user)

	// Return the newly created user
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
