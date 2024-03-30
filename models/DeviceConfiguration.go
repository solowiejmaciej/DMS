package models

import "gorm.io/gorm"

type DeviceConfiguration struct {
	gorm.Model
	DeviceId        string `gorm:"unique"`
	MqttBrokerPort  int
	MqttBrokerHost  string
	MqttUsername    string
	MqttPassword    string
	AliveInterval   int
	DeviceModel     string
	DeviceBoardType string
	SoftwareVersion string
	DmsKey          string
	DmsIV           string
	InternalLedPin  int
}
