package repositories

import (
	"dms/db"
	"dms/entities"
)

func AddDevice(device entities.Device) error {
	result := db.DB.Create(&device)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetByDeviceId(deviceId string) (entities.Device, error) {
	var device entities.Device
	result := db.DB.Where("device_id = ?", deviceId).First(&device)
	if result.Error != nil {
		return device, result.Error
	}
	return device, nil
}

func GetDevicesByUserId(userId uint) ([]entities.Device, error) {
	var devices []entities.Device
	result := db.DB.Where("user_id = ?", userId).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}
	return devices, nil
}

func GetDevices(userId uint) ([]entities.Device, error) {
	var devices []entities.Device
	result := db.DB.Where("user_id = ?", userId).Find(&devices)
	if result.Error != nil {
		return devices, result.Error
	}
	return devices, nil
}

func DeleteDevice(deviceId string, userId uint) error {
	result := db.DB.Where("device_id = ? AND user_id = ?", deviceId, userId).Delete(&entities.Device{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateDevice(device entities.Device, device2 entities.Device) {
	db.DB.Model(&device2).Updates(device)
}
