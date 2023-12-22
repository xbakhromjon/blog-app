package config

import (
	"github.com/joho/godotenv"
	"log"
)

func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
