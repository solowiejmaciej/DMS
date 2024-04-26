package models

type UserModel struct {
	ID                     uint   `json:"id"`
	FirstName              string `json:"firstname"`
	LastName               string `json:"lastname"`
	Email                  string `json:"email"`
	PhoneNumber            string `json:"phoneNumber"`
	DmsApiKey              string `json:"dmsApiKey"`
	Country                string `json:"country"`
	BirthDate              string `json:"birthDate"`
	IsPhoneNumberConfirmed bool   `json:"isPhoneNumberConfirmed"`
}
