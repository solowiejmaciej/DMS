package services

import (
	"encoding/base64"
	"github.com/google/uuid"
)

func GenerateApiKey() string {
	uuid := uuid.New()
	encodedUUID := base64.StdEncoding.EncodeToString([]byte(uuid.String()))
	return encodedUUID
}
