package entites

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	DeviceId string `gorm:"unique"`
	UserId   uint   `gorm:"not null"`
	User     User
}
