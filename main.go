package main

import (
	"log"

	"github.com/armanfarokhi/blog/database"
	"github.com/armanfarokhi/blog/handlers"
	"github.com/armanfarokhi/blog/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := database.InitDB()
	defer db.Close()

	router := gin.Default()

	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())

	router.Run(":8080")

	log.Println("Server is running on http://localhost:8080")
}
