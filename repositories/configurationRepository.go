package repositories

import (
	"dms/db"
	"dms/entites"
)

func AddConfiguration(deviceConfiguration entites.DeviceConfiguration) error {
	result := db.DB.Create(&deviceConfiguration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetConfigurationByDeviceId(deviceId string) (entites.DeviceConfiguration, error) {
	var device entites.DeviceConfiguration
	result := db.DB.Where("device_id = ?", deviceId).First(&device)
	if result.Error != nil {
		return device, result.Error
	}
	return device, nil
}

func UpdateConfiguration(deviceConfiguration entites.DeviceConfiguration) error {
	result := db.DB.Save(&deviceConfiguration)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteDeviceConfiguration(deviceId string) (int64, error) {
	result := db.DB.Where("device_id = ?", deviceId).Delete(&entites.DeviceConfiguration{})
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

func GetAllConfigurations() ([]entites.DeviceConfiguration, error) {
	var configurations []entites.DeviceConfiguration
	result := db.DB.Find(&configurations)
	if result.Error != nil {
		return configurations, result.Error
	}
	return configurations, nil
}
