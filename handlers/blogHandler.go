package handlers

import (
	"net/http"

	"github.com/armanfarokhi/blog/database"
	"github.com/armanfarokhi/blog/models"
	"github.com/gin-gonic/gin"
)

type CreateBlogInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdateBlogInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func CreateBlog(c *gin.Context) {
	var input CreateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ := c.Get("user")
	userID := user.(models.User).ID
	userEmail := user.(models.User).Email

	newBlog := models.Blog{
		Title:       input.Title,
		Content:     input.Content,
		AuthorID:    userID,
		AuthorEmail: userEmail,
	}

	if err := database.DB.Create(&newBlog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog post"})
		return
	}

	response := gin.H{
		"ID":          newBlog.ID,
		"Title":       newBlog.Title,
		"Content":     newBlog.Content,
		"AuthorID":    newBlog.AuthorID,
		"AuthorEmail": newBlog.AuthorEmail,
	}

	c.JSON(http.StatusCreated, response)
}

func GetBlogs(c *gin.Context) {
	var blogs []models.Blog

	if err := database.DB.
		Model(&models.Blog{}).
		Select("blogs.*, users.email as author_email").
		Joins("left join users on blogs.author_id = users.id").
		Where("blogs.deleted_at IS NULL").
		Scan(&blogs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blog posts"})
		return
	}

	var blogResponses []gin.H
	for _, blog := range blogs {
		blogResponse := gin.H{
			"ID":          blog.ID,
			"Title":       blog.Title,
			"Content":     blog.Content,
			"AuthorEmail": blog.AuthorEmail,
			"Likes":       blog.Likes,
		}
		blogResponses = append(blogResponses, blogResponse)
	}

	c.JSON(http.StatusOK, blogResponses)
}
func UpdateBlog(c *gin.Context) {

	blogID := c.Param("id")

	var blog models.Blog
	if err := database.DB.First(&blog, blogID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	user, _ := c.Get("user")
	userID := user.(models.User).ID
	if blog.AuthorID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update this blog"})
		return
	}

	var input UpdateBlogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blog.Title = input.Title
	blog.Content = input.Content

	if err := database.DB.Save(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func DeleteBlog(c *gin.Context) {

	blogID := c.Param("id")

	var blog models.Blog
	if err := database.DB.First(&blog, blogID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	user, _ := c.Get("user")
	userID := user.(models.User).ID
	if blog.AuthorID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to delete this blog"})
		return
	}

	if err := database.DB.Delete(&blog).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}
