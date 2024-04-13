package middleware

import (
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
)

func RequireDeviceOwnership(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	deviceId := c.Param("deviceId")
	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	res, err := services.IsDeviceOwnedByUser(userId, deviceId)
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if !res {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}
