package middleware

import (
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RequireDeviceToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(401, gin.H{"error": "Token is required"})
		c.Abort()
		return
	}
	var decryptedToken, err = services.Decrypt(token)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}
	log.Info(decryptedToken)
	_, err = repositories.GetByDeviceId(decryptedToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	deviceId := c.Param("deviceId")
	if deviceId != "" && decryptedToken != deviceId {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return

	}
	c.Next()
}
