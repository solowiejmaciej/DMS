package entites

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName              string
	LastName               string
	Email                  string `gorm:"unique"`
	Password               string
	PhoneNumber            string
	DmsApiKey              string `gorm:"unique index not null"`
	BirthDate              time.Time
	Country                string
	IsPhoneNumberConfirmed bool
}
