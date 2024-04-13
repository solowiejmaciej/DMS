package controllers

import (
	"dms/entites"
	"dms/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ActivateDevice(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	deviceId := c.Param("deviceId")

	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user details"})
		return
	}

	err = repositories.AddDevice(entites.Device{
		DeviceId: deviceId,
		UserId:   userId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while adding device"})
		log.Error("Error while adding device", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Device activated successfully"})
}

func DeleteDevice(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	deviceId := c.Param("deviceId")

	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user details"})
		return
	}

	_, err = repositories.DeleteDeviceConfiguration(deviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while deleting device configuration"})
		log.Error("Error while deleting device configuration", err)
		return
	}

	_, err = repositories.GetByDeviceId(deviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found"})
		return
	}

	err = repositories.DeleteDevice(deviceId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while deleting device"})
		log.Error("Error while deleting device", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Device deleted successfully"})
}
