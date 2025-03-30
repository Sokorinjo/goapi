package main

import (
	// "fmt"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Penis"))
	})

	router.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page"))
	})

	router.Get("/users/{userId}-{userName}", getUser)

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Printf("Error occured: %v", err)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	userName := chi.URLParam(r, "userName")

	w.Write([]byte("User: " + userId + " " + userName))
}
