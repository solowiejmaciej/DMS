package initializers

import (
	log "github.com/sirupsen/logrus"
	"notificator/entities"
)

func Migrate() {
	notificationError := DB.AutoMigrate(&entities.Notification{})
	if notificationError != nil {
		log.Error("Error while performing migration")
		panic(notificationError)
	}
}
