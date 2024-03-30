package controllers

import (
	"dms/initializers"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RebootDevice(c *gin.Context) {
	deviceId := c.Param("deviceId")

	_, err3 := repositories.GetByDeviceId(deviceId)
	if err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found"})
		return
	}

	services.PublishMessageToTopic(initializers.Client, "/api/"+deviceId+"/reboot", "{}")

	c.JSON(http.StatusOK, gin.H{"message": "Device rebooted successfully"})
}

func GetDevices(c *gin.Context) {
	devices, err := repositories.GetDevices()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Devices not found"})
		return
	}
	c.JSON(http.StatusOK, devices)
}

func GetDevice(c *gin.Context) {
	deviceId := c.Param("deviceId")
	device, err := repositories.GetByDeviceId(deviceId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}
	c.JSON(http.StatusOK, device)
}
