package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Users Page
func usersPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users Page"))
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
