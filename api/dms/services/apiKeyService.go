package services

import (
	"dms/repositories"
	"encoding/base64"
	"github.com/google/uuid"
)

func GenerateApiKey() string {
	uuid := uuid.New()
	encodedUUID := base64.StdEncoding.EncodeToString([]byte(uuid.String()))
	return encodedUUID
}

func RegenerateApiKeyForUser(userId uint) (string, error) {
	user, err := repositories.GetById(userId)
	if err != nil {
		return "", err
	}
	user.DmsApiKey = GenerateApiKey()
	updatedUser, err := repositories.UpdateUser(user)
	if err != nil {
		return "", err
	}
	return updatedUser.DmsApiKey, nil
}
