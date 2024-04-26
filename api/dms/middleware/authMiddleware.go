package middleware

import (
	"dms/managers"
	"github.com/gin-gonic/gin"
)

func RequireToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Token is required"})
		c.Abort()
		return
	}
	var isValidToken = managers.ValidateToken(token)
	if !isValidToken {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	c.Next()
}
