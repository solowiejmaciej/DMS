package controllers

import (
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
)

func GetConfiguration(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var decryptedToken, _ = services.Decrypt(token)
	device, _ := repositories.GetByDeviceId(decryptedToken)
	var deviceConfig, err = repositories.GetConfigurationByDeviceId(device.DeviceId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Configuration not found"})
		return
	}

	c.JSON(200, deviceConfig)
}

func AddConfiguration(c *gin.Context) {
	token := c.GetHeader("Authorization")
	var decryptedToken, _ = services.Decrypt(token)
	device, _ := repositories.GetByDeviceId(decryptedToken)
	var deviceConfigBody struct {
		MqttBrokerPort int    `json:"mqqt_broker_port"`
		MqttBrokerHost string `json:"mqqt_broker_host"`
		MqttUsername   string `json:"mqqt_username"`
		MqttPassword   string `json:"mqqt_password"`
		AliveInterval  int    `json:"alive_interval"`
	}
	err := c.ShouldBindJSON(&deviceConfigBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	var configuration = models.DeviceConfiguration{
		DeviceId:       device.DeviceId,
		MqttBrokerPort: deviceConfigBody.MqttBrokerPort,
		MqttBrokerHost: deviceConfigBody.MqttBrokerHost,
		MqttUsername:   deviceConfigBody.MqttUsername,
		MqttPassword:   deviceConfigBody.MqttPassword,
		AliveInterval:  deviceConfigBody.AliveInterval,
	}
	err = repositories.AddConfiguration(configuration)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while adding configuration"})
		return
	}

	c.JSON(200, gin.H{"status": "ok"})
}
