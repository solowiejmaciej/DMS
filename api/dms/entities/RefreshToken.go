package entities

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	UserId       uint `gorm:"not null"`
	User         User
	RefreshToken string `gorm:"unique;not null"`
	ExpiresAt    int64  `gorm:"not null"`
	IsValid      bool   `gorm:"not null"`
	Jit          string `gorm:"not null"`
}
