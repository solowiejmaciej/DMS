package repositories

import (
	"dms/db"
	"dms/entities"
)

func AddConfiguration(deviceConfiguration entities.DeviceConfiguration) error {
	result := db.DB.Create(&deviceConfiguration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetConfigurationByDeviceId(deviceId string) (entities.DeviceConfiguration, error) {
	var device entities.DeviceConfiguration
	result := db.DB.Where("device_id = ?", deviceId).First(&device)
	if result.Error != nil {
		return device, result.Error
	}
	return device, nil
}

func UpdateConfiguration(deviceConfiguration entities.DeviceConfiguration) error {
	result := db.DB.Save(&deviceConfiguration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteDeviceConfiguration(deviceId string) (int64, error) {
	result := db.DB.Where("device_id = ?", deviceId).Delete(&entities.DeviceConfiguration{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func GetAllConfigurations() ([]entities.DeviceConfiguration, error) {
	var configurations []entities.DeviceConfiguration
	result := db.DB.Find(&configurations)
	if result.Error != nil {
		return configurations, result.Error
	}
	return configurations, nil
}
