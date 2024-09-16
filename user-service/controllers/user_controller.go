package controllers

import (
	"encoding/json"
	"fmt"
	"go-microservices-app/user-service/models"
	"net/http"

	"github.com/gorilla/mux"
)

var users []models.User // Slice to hold user data

func init() {
	// Seed with some dummy users
	users = append(users, models.User{UserID: "1", UserName: "John Doe", UserEmail: "john@example.com"})
	users = append(users, models.User{UserID: "2", UserName: "Jane Smith", UserEmail: "jane@example.com"})
}

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello this is the user page guys!</h1>"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, user := range users {
		if user.UserID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating one user....!")
	w.Header().Set("Content-Type", "application/json")

	//what if the body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Loop through users to check for duplicate product
	for _, existingUser := range users {
		if existingUser.UserName == user.UserName {
			json.NewEncoder(w).Encode("user name already exists")
			return
		}
	}

	//append product into products

	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

//update users

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedUser models.User
	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	for index, user := range users {
		if user.UserID == params["id"] {
			// Update the user in the slice
			users[index] = updatedUser

			// Return the updated user
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

//deleting the user

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for index, user := range users {
		if user.UserID == params["id"] {
			// Remove the user from the slice
			users = append(users[:index], users[index+1:]...)

			// Return a success message
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}
