package controllers

import (
	"dms/entities"
	"dms/initializers"
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
)

func GetConfiguration(c *gin.Context) {
	deviceId := c.Param("deviceId")
	deviceConfig, err := repositories.GetConfigurationByDeviceId(deviceId)

	model := models.GetDeviceConfigurationModel{
		MqttBrokerPort: deviceConfig.MqttBrokerPort,
		MqttBrokerHost: deviceConfig.MqttBrokerHost,
		MqttUsername:   deviceConfig.MqttUsername,
		MqttPassword:   deviceConfig.MqttPassword,
		AliveInterval:  deviceConfig.AliveInterval,
		InternalLedPin: deviceConfig.InternalLedPin,
		CreatedAt:      deviceConfig.CreatedAt.Format("2006-01-02 15:04"),
	}

	if err != nil {
		c.JSON(404, gin.H{"error": "Configuration not found"})
		return

	}

	c.JSON(200, model)
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
	var configuration = entities.DeviceConfiguration{
		DeviceId:       device.DeviceId,
		MqttBrokerPort: deviceConfigurationModel.MqttBrokerPort,
		MqttBrokerHost: deviceConfigurationModel.MqttBrokerHost,
		MqttUsername:   deviceConfigurationModel.MqttUsername,
		MqttPassword:   deviceConfigurationModel.MqttPassword,
		AliveInterval:  deviceConfigurationModel.AliveInterval,
		InternalLedPin: deviceConfigurationModel.InternalLedPin,
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

func UpdateOrCreateConfiguration(c *gin.Context) {
	var deviceConfigurationModel models.DeviceConfigurationModel
	err := c.ShouldBindJSON(&deviceConfigurationModel)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	deviceId := c.Param("deviceId")
	deviceConfig, err := repositories.GetConfigurationByDeviceId(deviceId)

	if err != nil {
		device, _ := repositories.GetByDeviceId(deviceId)
		if device.DeviceId == "" {
			c.JSON(404, gin.H{"error": "Device not found"})
			return
		}

		var configuration = entities.DeviceConfiguration{
			DeviceId:       device.DeviceId,
			MqttBrokerPort: deviceConfigurationModel.MqttBrokerPort,
			MqttBrokerHost: deviceConfigurationModel.MqttBrokerHost,
			MqttUsername:   deviceConfigurationModel.MqttUsername,
			MqttPassword:   deviceConfigurationModel.MqttPassword,
			AliveInterval:  deviceConfigurationModel.AliveInterval,
			InternalLedPin: deviceConfigurationModel.InternalLedPin,
		}

		err = repositories.AddConfiguration(configuration)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error while adding configuration"})
			return
		}
		services.ConfigureMqttForDevice(deviceConfig)

		c.JSON(200, gin.H{"status": "ok"})
	}

	deviceConfig.MqttBrokerPort = deviceConfigurationModel.MqttBrokerPort
	deviceConfig.MqttBrokerHost = deviceConfigurationModel.MqttBrokerHost
	deviceConfig.MqttUsername = deviceConfigurationModel.MqttUsername
	deviceConfig.MqttPassword = deviceConfigurationModel.MqttPassword
	deviceConfig.AliveInterval = deviceConfigurationModel.AliveInterval
	deviceConfig.InternalLedPin = deviceConfigurationModel.InternalLedPin

	err = repositories.UpdateConfiguration(deviceConfig)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error while updating configuration"})
		return
	}
	services.ConfigureMqttForDevice(deviceConfig)
	c.JSON(200, gin.H{"status": "ok"})
}
