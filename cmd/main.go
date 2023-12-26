package main

import (
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/config"
	"golang-project-template/internal/delivery/rest/router"
	"log"
	"net/http"
	"os"
)

func main() {

	config.SetupEnv()
	db := config.SetupDB()
	defer db.Close()
	config.SetupIdentityProviders()
	r := chi.NewRouter()
	router.SetupRouter(r, db)
	err := http.ListenAndServe(os.Getenv("HTTP_PORT"), r)
	if err != nil {
		log.Fatal(err)
	}
}
