package services

import "dms/repositories"

func IsDeviceOwnedByUser(userId uint, deviceId string) (bool, error) {
	device, err := repositories.GetByDeviceId(deviceId)
	if err != nil {
		return false, err
	}

	return device.UserId == userId, nil
}
