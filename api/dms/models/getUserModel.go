package models

import "time"

type GetUserModel struct {
	Id          uint      `json:"id"`
	FirstName   string    `json:"firstname"`
	LastName    string    `json:"lastname"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phoneNumber"`
	CreatedAt   time.Time `json:"createdAt"`
	Country     string    `json:"country"`
	BirthDate   string    `json:"birthDate"`
}
