package repositories

import (
	"dms/db"
	"dms/entities"
)

func SaveRefreshToken(token entities.RefreshToken) error {
	return db.DB.Create(&token).Error
}

func GetRefreshToken(token string) (entities.RefreshToken, error) {
	var refreshToken entities.RefreshToken
	result := db.DB.Where("refresh_token = ?", token).First(&refreshToken)
	if result.Error != nil {
		return refreshToken, result.Error
	}
	return refreshToken, nil
}

func InvalidateTokenByValue(token string) {
	refreshToken, _ := GetRefreshToken(token)
	refreshToken.IsValid = false
	db.DB.Save(refreshToken)
}
