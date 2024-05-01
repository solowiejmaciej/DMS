package repositories

import (
	log "github.com/sirupsen/logrus"
	"notificator/initializers"
)

func GetUserPhoneNumber(userId uint) (string, error) {
	var phoneNumber string
	result := initializers.DB.Raw("SELECT phone_number FROM users WHERE id = ?", userId).Scan(&phoneNumber)
	if result.Error != nil {
		log.Errorf("Error while executing raw SQL: %v", result.Error)
		return "", result.Error
	}
	return phoneNumber, nil
}
