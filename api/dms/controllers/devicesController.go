package controllers

import (
	"dms/models"
	"dms/repositories"
	"dms/services"
	"encoding/json"
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

	services.PublishMessageToTopic(deviceId, "/api/"+deviceId+"/reboot", "{}")

	c.JSON(http.StatusOK, gin.H{"message": "Reboot message sent"})
}

func GetDevices(c *gin.Context) {
	apiKey := c.GetHeader("X-Api-Key")
	err, userId := repositories.GetByApiKey(apiKey)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	devices, err := repositories.GetDevices(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Devices not found"})
		return
	}

	devicesResponse := make([]models.DeviceModel, 0)
	for _, device := range devices {
		devicesResponse = append(devicesResponse, models.DeviceModel{
			DeviceId:        device.DeviceId,
			UserId:          device.UserId,
			CreatedAt:       device.CreatedAt.Format("2006-01-02 15:04"),
			SoftwareVersion: device.SoftwareVersion,
			DeviceModel:     device.DeviceModel,
			DeviceBoardType: device.DeviceBoardType,
			CodeBaseUrl:     device.CodeBaseUrl,
		})
	}
	c.JSON(http.StatusOK, devicesResponse)
}

func GetDeviceStatus(c *gin.Context) {
	deviceId := c.Param("deviceId")
	status := services.CalculateDeviceStatus(deviceId)
	c.JSON(http.StatusOK, gin.H{"status": status})
}

func GetDevice(c *gin.Context) {
	deviceId := c.Param("deviceId")
	device, err := repositories.GetByDeviceId(deviceId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device not found"})
		return
	}

	c.JSON(http.StatusOK, models.DeviceModel{
		DeviceId:        device.DeviceId,
		UserId:          device.UserId,
		CreatedAt:       device.CreatedAt.Format("2006-01-02 15:04"),
		SoftwareVersion: device.SoftwareVersion,
		DeviceModel:     device.DeviceModel,
		DeviceBoardType: device.DeviceBoardType,
		CodeBaseUrl:     device.CodeBaseUrl,
	})
}

func UpdateDeviceFirmware(c *gin.Context) {
	deviceId := c.Param("deviceId")

	var firmwareBody struct {
		SoftwareVersion string `json:"software_version"`
	}

	var err = c.ShouldBindJSON(&firmwareBody)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	_, err = repositories.GetByDeviceId(deviceId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found"})
		return
	}

	messageBody, err := json.Marshal(firmwareBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while marshalling firmware body"})
		return

	}

	services.PublishMessageToTopic(deviceId, "/api/"+deviceId+"/update", string(messageBody))

	c.JSON(http.StatusOK, gin.H{"message": "Update message sent"})
}

func TestMqttConnection(context *gin.Context) {
	var body struct {
		MqttBrokerHost string `json:"mqtt_broker_host"`
		MqttBrokerPort int    `json:"mqtt_broker_port"`
		MqttUsername   string `json:"mqtt_username"`
		MqttPassword   string `json:"mqtt_password"`
	}
	err := context.ShouldBindJSON(&body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, connectionError := services.TestMqttConnection(body.MqttBrokerHost, body.MqttBrokerPort, body.MqttUsername, body.MqttPassword)

	if connectionError != nil {
		context.JSON(http.StatusOK, gin.H{"result": res, "error": connectionError.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"result": res, "error": nil})

}
