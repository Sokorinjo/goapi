package user

import (
	"github.com/go-chi/chi/v5"
)

func UserRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getAllUsers)
	r.Post("/", createUser)
	r.Route("/{userId}", func(r chi.Router) {
		r.Get("/", getUser)
		r.Delete("/", deleteUser)
	})

	return r
}
