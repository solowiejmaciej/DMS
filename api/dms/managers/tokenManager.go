package managers

import (
	"dms/entites"
	"dms/models"
	"dms/repositories"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func GenerateToken(user entites.User) (error, string) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"userId": user.ID,
		"iss":    "DMS",
		"exp":    time.Now().Add(time.Hour).Unix(),
		"iat":    time.Now().Unix(),
		"nbf":    time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Error("Error while generating token", err)
		return err, ""
	}
	return nil, tokenString
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if time.Now().Unix() > int64(claims["exp"].(float64)) {
			return false
		}
	}
	return token.Valid
}

func GetUserFromToken(token string) (error, models.UserModel) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err, models.UserModel{}
	}

	var userId = uint(claims["userId"].(float64))
	var user, userError = repositories.GetById(userId)
	if userError != nil {
		return err, models.UserModel{}
	}

	return nil, models.UserModel{
		ID:                     user.ID,
		FirstName:              user.FirstName,
		LastName:               user.LastName,
		Email:                  user.Email,
		PhoneNumber:            user.PhoneNumber,
		DmsApiKey:              user.DmsApiKey,
		Country:                user.Country,
		BirthDate:              user.BirthDate.Format("2006-01-02"),
		IsPhoneNumberConfirmed: user.IsPhoneNumberConfirmed,
	}

}
