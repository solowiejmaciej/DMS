package repositories

import (
	"dms/db"
	"dms/entities"
)

func AddUser(User entities.User) (entities.User, error) {
	result := db.DB.Create(&User)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return User, nil
}

func GetById(id uint) (entities.User, error) {
	var user entities.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GetByApiKey(apiKey string) (error error, userId uint) {
	var user entities.User
	result := db.DB.Where("dms_api_key = ?", apiKey).First(&user)
	if result.Error != nil {
		return result.Error, user.ID
	}
	return nil, user.ID
}

func GetByEmail(email string) (error, entities.User) {
	var user entities.User
	result := db.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return result.Error, user
	}
	return nil, user
}

func UpdateUser(user entities.User) (entities.User, error) {
	result := db.DB.Save(&user)
	if result.Error != nil {
		return entities.User{}, result.Error
	}
	return user, nil
}
