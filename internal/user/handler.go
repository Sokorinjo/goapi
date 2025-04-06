package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Get all users
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users."))
}

// Get specific user
func getUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	w.Write([]byte("User id: " + userId))
}

// Create user
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}
