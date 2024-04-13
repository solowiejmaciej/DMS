package initializers

import (
	"dms/db"
	"dms/entites"
	log "github.com/sirupsen/logrus"
)

func Migrate() {
	err := db.DB.AutoMigrate(&entites.Device{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}

	err = db.DB.AutoMigrate(&entites.DeviceConfiguration{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}

	err = db.DB.AutoMigrate(&entites.User{})
	if err != nil {
		log.Error("Error while performing migration")
		panic(err)
	}
}
