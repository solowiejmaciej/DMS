package repositories

import (
	"dms/db"
	"dms/entites"
)

func AddUser(User entites.User) (entites.User, error) {
	result := db.DB.Create(&User)
	if result.Error != nil {
		return entites.User{}, result.Error
	}
	return User, nil
}

func GetById(id uint) (entites.User, error) {
	var user entites.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
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

func GetByEmail(email string) (error, entites.User) {
	var user entites.User
	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return result.Error, user
	}
	return nil, user
}

func UpdateUser(user entites.User) (entites.User, error) {
	result := db.DB.Save(&user)
	if result.Error != nil {
		return entites.User{}, result.Error
	}
	return user, nil
}
