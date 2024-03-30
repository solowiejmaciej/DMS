package repositories

import (
	"dms/initializers"
	"dms/models"
)

func AddConfiguration(deviceConfiguration models.DeviceConfiguration) error {
	result := initializers.DB.Create(&deviceConfiguration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetConfigurationByDeviceId(deviceId string) (models.DeviceConfiguration, error) {
	var device models.DeviceConfiguration
	result := initializers.DB.Where("device_id = ?", deviceId).First(&device)
	if result.Error != nil {
		return device, result.Error
	}
	return device, nil
}
