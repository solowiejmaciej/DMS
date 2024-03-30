package models

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	DeviceId string `gorm:"unique"`
	Token    string
}
