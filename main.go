package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"

	"github.com/Sokorinjo/goapi/internal/db"
	"github.com/Sokorinjo/goapi/internal/post"
	"github.com/Sokorinjo/goapi/internal/user"
)

func main() {

	loadEnvErr := godotenv.Load()
	if loadEnvErr != nil {
		log.Fatal("Failed to load env.")
	}

	db.ConnectDB()
	defer db.CloseDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page"))
	})

	r.Mount("/users", user.UserRoutes())
	r.Mount("/posts", post.PostRoutes())

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("Route does not exist."))
	})

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		fmt.Println(err)
	}
}
