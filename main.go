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

	database.InitDB()
	defer database.DB.Close()

	router := gin.Default()

	router.POST("/signup", handlers.Signup)
	router.POST("/login", handlers.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/blog", handlers.CreateBlog)
		authorized.GET("/blog", handlers.GetBlogs)
		authorized.PUT("/blog/:id", handlers.UpdateBlog)
		authorized.DELETE("/blog/:id", handlers.DeleteBlog)
	}

	authorized.POST("/like", handlers.LikeBlog)
	authorized.DELETE("/like", handlers.UnlikeBlog)

	log.Println("Server running on port 8080")
	router.Run(":8080")
}
