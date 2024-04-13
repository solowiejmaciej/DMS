package repositories

import (
	"dms/db"
	"dms/entites"
)

func AddDevice(device entites.Device) error {
	result := db.DB.Create(&device)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetByDeviceId(deviceId string) (entites.Device, error) {
	var device entites.Device
	result := db.DB.Where("device_id = ?", deviceId).First(&device)
	if result.Error != nil {
		return device, result.Error
	}
	return device, nil
}

func GetDevicesByUserId(userId uint) ([]entites.Device, error) {
	var devices []entites.Device
	result := db.DB.Where("user_id = ?", userId).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}
	return devices, nil
}

func GetDevices(userId uint) ([]entites.Device, error) {
	var devices []entites.Device
	result := db.DB.Where("user_id = ?", userId).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}
	return devices, nil
}

func DeleteDevice(deviceId string, userId uint) error {
	result := db.DB.Where("device_id = ? AND user_id = ?", deviceId, userId).Delete(&entites.Device{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
