package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement authentication logic here
		// Example: Check JWT token validity

		// Skip middleware for now (replace with actual authentication logic)
		c.Next()
	}
}
