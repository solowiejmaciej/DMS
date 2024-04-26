package entites

import "gorm.io/gorm"

type DeviceConfiguration struct {
	gorm.Model
	DeviceId       string `gorm:"unique"`
	MqttBrokerPort int    `gorm:"not null"`
	MqttBrokerHost string `gorm:"not null"`
	MqttUsername   string `gorm:"not null"`
	MqttPassword   string `gorm:"not null"`
	AliveInterval  int    `gorm:"not null"`
	InternalLedPin int    `gorm:"not null"`
}
