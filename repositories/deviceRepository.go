package repositories

import (
	"dms/initializers"
	"dms/models"
)

func AddDevice(device models.Device) error {
	result := initializers.DB.Create(&device)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetByDeviceId(deviceId string) (models.Device, error) {
	var device models.Device
	result := initializers.DB.Where("device_id = ?", deviceId).First(&device)
	if result.Error != nil {
		return device, result.Error
	}
	return device, nil
}

func GetDevices() ([]models.Device, error) {
	var devices []models.Device
	result := initializers.DB.Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}
	return devices, nil
}
