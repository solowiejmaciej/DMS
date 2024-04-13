package repositories

import (
	"dms/db"
	"dms/entites"
)

func AddUser(User entites.User) (uint, error) {
	result := db.DB.Create(&User)
	if result.Error != nil {
		return 0, result.Error
	}
	return User.ID, nil
}

func GetById(id int) (entites.User, error) {
	var user entites.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	user.Password = ""
	return user, nil
}

func GetByApiKey(apiKey string) (error error, userId uint) {
	var user entites.User
	result := db.DB.Where("dms_api_key = ?", apiKey).First(&user)
	if result.Error != nil {
		return result.Error, user.ID
	}
	return nil, user.ID
}
