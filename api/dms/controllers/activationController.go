package controllers

import (
	"dms/entities"
	"dms/models"
	"dms/repositories"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func ActivateDevice(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	deviceId := c.Param("deviceId")

	var body models.ActivateDeviceModel
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return

	}
	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while fetching user details"})
		return
	}

	currentDevice, _ := repositories.GetByDeviceId(deviceId)

	if currentDevice.DeviceId != "" {
		if userId == currentDevice.UserId {
			repositories.UpdateDevice(entities.Device{
				DeviceId:        deviceId,
				DeviceModel:     body.DeviceModel,
				DeviceBoardType: body.DeviceBoardType,
				SoftwareVersion: body.SoftwareVersion,
				CodeBaseUrl:     body.CodeBaseUrl,
			}, currentDevice)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Another user is using this device. Please delete it first"})
			return
		}
	} else {
		err = repositories.AddDevice(entities.Device{
			DeviceId:        deviceId,
			UserId:          userId,
			DeviceModel:     body.DeviceModel,
			DeviceBoardType: body.DeviceBoardType,
			SoftwareVersion: body.SoftwareVersion,
			CodeBaseUrl:     body.CodeBaseUrl,
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error while adding device"})
			log.Error("Error while adding device", err)
			return
		}

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
