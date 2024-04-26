package models

type AddUserModel struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
	Country     string `json:"country"`
	BirthDate   string `json:"birthDate"`
}
