package handlers

import (
	"errors"
	"net/http"

	"github.com/armanfarokhi/blog/database"
	"github.com/armanfarokhi/blog/models"
	"github.com/gin-gonic/gin"
)

func getUserIdFromToken(c *gin.Context) (uint, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, errors.New("Authorization header required")
	}

	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}

	return claims.UserID, nil
}

func LikeBlog(c *gin.Context) {
	userID, err := getUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var userLike models.UserLikeBlog
	if err := database.DB.Where("user_id = ? AND blog_id = ?", userID, input.ID).First(&userLike).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have already liked this post"})
		return
	}

	newLike := models.UserLikeBlog{
		UserID: userID,
		BlogID: input.ID,
	}
	if err := database.DB.Create(&newLike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like the post"})
		return
	}

	if err := updateBlogLikesCount(input.ID, 1); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog likes count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}

func UnlikeBlog(c *gin.Context) {
	userID, err := getUserIdFromToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var input struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var userLike models.UserLikeBlog
	if err := database.DB.Where("user_id = ? AND blog_id = ?", userID, input.ID).First(&userLike).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You have not liked this post"})
		return
	}

	if err := database.DB.Delete(&userLike).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unlike the post"})
		return
	}

	if err := updateBlogLikesCount(input.ID, -1); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog likes count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
}

func updateBlogLikesCount(blogID uint, increment int) error {
	var blog models.Blog
	if err := database.DB.Where("id = ?", blogID).First(&blog).Error; err != nil {
		return err
	}

	blog.Likes += increment
	if err := database.DB.Save(&blog).Error; err != nil {
		return err
	}

	return nil
}
