package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/create-user-test-case/models"
	"github.com/gorilla/mux"
)

// CreateUser handles HTTP POST requests to create a user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	// Here you would normally save the user to a database and generate an ID
	user.ID = 1 // Example: In real scenario, generate ID from DB

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// SetupRoutes sets up the routes for the application
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", CreateUser).Methods("POST")
	return r
}
