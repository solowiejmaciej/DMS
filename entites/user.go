package entites

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string `gorm:"unique"`
	Password    string
	PhoneNumber string
	DmsApiKey   string `gorm:"unique index not null"`
}
