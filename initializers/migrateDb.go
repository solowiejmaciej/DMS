package initializers

import (
	"dms/models"
	log "github.com/sirupsen/logrus"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Device{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}

	err = DB.AutoMigrate(&models.DeviceConfiguration{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}
}
