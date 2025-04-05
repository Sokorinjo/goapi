package post

import (
	"github.com/go-chi/chi/v5"
)

func PostRoutes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", getPosts)
	r.Post("/", createPost)

	return r
}
