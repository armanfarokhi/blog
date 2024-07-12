package main

import (
	"log"

	"github.com/armanfarokhi/blog/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := database.InitDB
	defer db.Close()

}
