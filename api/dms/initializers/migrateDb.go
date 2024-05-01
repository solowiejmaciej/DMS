package initializers

import (
	"dms/db"
	"dms/entities"
	log "github.com/sirupsen/logrus"
)

func Migrate() {
	log.Info("Performing migration")
	err := db.DB.AutoMigrate(&entities.Device{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}

	err = db.DB.AutoMigrate(&entities.DeviceConfiguration{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}

	err = db.DB.AutoMigrate(&entities.User{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}

	err = db.DB.AutoMigrate(&entities.RefreshToken{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}
}
