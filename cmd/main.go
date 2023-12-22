package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang-project-template/internal/config"
	"log"
	"net/http"
	"os"
)

func main() {
	config.SetupEnv()
	config.SetupDB()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	err := http.ListenAndServe(os.Getenv("HTTP_PORT"), r)
	if err != nil {
		log.Fatal(err)
	}
}
