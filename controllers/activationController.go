package controllers

import (
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ActivateDevice(c *gin.Context) {
	var body struct {
		DeviceToken string `json:"device_token"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var decryptedToken, err2 = services.Decrypt(body.DeviceToken)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device id"})
		log.Error("Error while decrypting token", err2)
		return
	}

	err3 := repositories.AddDevice(models.Device{
		DeviceId: decryptedToken,
		Token:    body.DeviceToken,
	})
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while adding device"})
		log.Error("Error while adding device", err2)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Device activated successfully"})
}
