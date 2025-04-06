package user

import (
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllUsers)
	r.Post("/", createUser)
	r.Get("/{userId}", getUser)

	return r
}