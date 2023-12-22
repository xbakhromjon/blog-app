package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func SetupDB() {
	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DATABASE"),
	)
	db, _ := sql.Open("postgres", connString)

	defer db.Close()
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected db")
}
