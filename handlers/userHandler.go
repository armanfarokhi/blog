package handlers

import (
	"log"
	"net/http"

	"github.com/armanfarokhi/blog/database"
	"github.com/armanfarokhi/blog/models"
	"github.com/armanfarokhi/blog/utils"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		log.Fatal(err)
	}
	newUser.Password = hashedPassword

	db := database.DB
	db.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfuly", "user": newUser})
}
