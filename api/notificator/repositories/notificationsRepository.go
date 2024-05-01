package repositories

import (
	"notificator/entities"
	"notificator/initializers"
)

func AddNewNotification(notification entities.Notification) (uint, error) {
	result := initializers.DB.Create(&notification)
	if result.Error != nil {
		return 0, result.Error
	}
	return notification.ID, nil
}

func GetNotificationsByUserId(userId int) ([]entities.Notification, error) {
	var notifications []entities.Notification
	result := initializers.DB.Where("user_id = ?", userId).Find(&notifications)
	if result.Error != nil {
		return notifications, result.Error
	}
	return notifications, nil
}
