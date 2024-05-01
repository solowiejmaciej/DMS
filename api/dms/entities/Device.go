package entities

import "gorm.io/gorm"

type Device struct {
	gorm.Model
	DeviceId        string `gorm:"unique"`
	UserId          uint   `gorm:"not null"`
	SoftwareVersion string `gorm:"not null"`
	DeviceModel     string `gorm:"not null"`
	DeviceBoardType string `gorm:"not null"`
	CodeBaseUrl     string
	User            User
}
