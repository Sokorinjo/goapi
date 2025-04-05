package post

import (
	"net/http"
	// "github.com/go-chi/chi/v5"
)

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All Posts"))
}

func createPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Post"))
}
