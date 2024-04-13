package controllers

import (
	"dms/entites"
	"dms/initializers"
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
)

func GetConfiguration(c *gin.Context) {
	deviceId := c.Param("deviceId")
	deviceConfig, err := repositories.GetConfigurationByDeviceId(deviceId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Configuration not found"})
		return

	}

	c.JSON(200, deviceConfig)
}

func AddConfiguration(c *gin.Context) {
	deviceId := c.Param("deviceId")
	device, _ := repositories.GetByDeviceId(deviceId)
	if device.DeviceId == "" {
		c.JSON(404, gin.H{"error": "Device not found"})
		return
	}
	var deviceConfigurationModel models.DeviceConfigurationModel
	err := c.ShouldBindJSON(&deviceConfigurationModel)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	var configuration = entites.DeviceConfiguration{
		DeviceId:        device.DeviceId,
		MqttBrokerPort:  deviceConfigurationModel.MqttBrokerPort,
		MqttBrokerHost:  deviceConfigurationModel.MqttBrokerHost,
		MqttUsername:    deviceConfigurationModel.MqttUsername,
		MqttPassword:    deviceConfigurationModel.MqttPassword,
		AliveInterval:   deviceConfigurationModel.AliveInterval,
		DeviceModel:     deviceConfigurationModel.DeviceModel,
		DeviceBoardType: deviceConfigurationModel.DeviceBoardType,
		InternalLedPin:  deviceConfigurationModel.InternalLedPin,
	}
	isAnyConfiguration, _ := repositories.GetConfigurationByDeviceId(device.DeviceId)
	if isAnyConfiguration.DeviceId != "" {
		c.JSON(400, gin.H{"error": "Configuration already exists"})
		return
	}
	err = repositories.AddConfiguration(configuration)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error while adding configuration"})
		return
	}

	initializers.RedisClient.Del(c.Request.URL.RequestURI())
	services.ConfigureMqttForDevice(configuration)
	c.JSON(200, gin.H{"status": "ok"})
}

func UpdateConfiguration(c *gin.Context) {
	deviceId := c.Param("deviceId")
	deviceConfig, err := repositories.GetConfigurationByDeviceId(deviceId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Configuration not found"})
		return
	}
	var deviceConfigurationModel models.DeviceConfigurationModel
	err = c.ShouldBindJSON(&deviceConfigurationModel)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	deviceConfig.MqttBrokerPort = deviceConfigurationModel.MqttBrokerPort
	deviceConfig.MqttBrokerHost = deviceConfigurationModel.MqttBrokerHost
	deviceConfig.MqttUsername = deviceConfigurationModel.MqttUsername
	deviceConfig.MqttPassword = deviceConfigurationModel.MqttPassword
	deviceConfig.AliveInterval = deviceConfigurationModel.AliveInterval
	deviceConfig.DeviceModel = deviceConfigurationModel.DeviceModel
	deviceConfig.DeviceBoardType = deviceConfigurationModel.DeviceBoardType
	deviceConfig.InternalLedPin = deviceConfigurationModel.InternalLedPin

	err = repositories.UpdateConfiguration(deviceConfig)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error while updating configuration"})
		return
	}
	services.ConfigureMqttForDevice(deviceConfig)
	c.JSON(200, gin.H{"status": "ok"})
}
