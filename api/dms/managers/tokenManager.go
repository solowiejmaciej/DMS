package managers

import (
	"dms/entities"
	"dms/models"
	"dms/repositories"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func GenerateToken(user entities.User) (error, string) {
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  user.Email,
		"userId": user.ID,
		"iss":    "DMS",
		"exp":    time.Now().Add(time.Second * 15).Unix(),
		"iat":    time.Now().Unix(),
		"nbf":    time.Now().Unix(),
		"jit":    uuid.New().String(),
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

func GetUserFromValidToken(token string) (error, models.UserModel) {
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

func GenerateRefreshToken(token string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		log.Error("Error while parsing token", err)
		return "", err
	}

	var userId = uint(claims["userId"].(float64))
	var jit = claims["jit"].(string)

	var refreshTokenString = uuid.New().String()

	refreshToken := entities.RefreshToken{
		UserId:       userId,
		RefreshToken: refreshTokenString,
		Jit:          jit,
		ExpiresAt:    time.Now().Add(time.Hour * 24 * 30).Unix(),
		IsValid:      true,
	}

	err = repositories.SaveRefreshToken(refreshToken)

	if err != nil {
		log.Error("Error while saving refresh token", err)
		return "", err
	}

	return refreshTokenString, nil
}

func RefreshToken(token string, refreshToken string) (string, error) {
	var userId, err = GetUserIdFromToken(token)
	if err != nil {
		return "", err
	}

	jit, jitErr := getJitFromToken(token)
	if jitErr != nil {
		log.Errorf("Unable to get jit from token: %v", err)
		return "", err
	}

	refreshTokenEntity, err := repositories.GetRefreshToken(refreshToken)
	if err != nil {
		return "", err
	}
	if refreshTokenEntity.ExpiresAt <= time.Now().Unix() {
		return "", errors.New("Token is not valid")
	}

	if !refreshTokenEntity.IsValid {
		return "", errors.New("RefreshToken is not valid")
	}

	if refreshTokenEntity.UserId != userId {
		return "", errors.New("RefreshToken not valid for this user")
	}

	if refreshTokenEntity.Jit != jit {
		return "", errors.New("refresh token not valid for this jwt")
	}

	userEntity, userErr := repositories.GetById(userId)

	if userErr != nil {
		return "", errors.New("error while getting user")
	}

	jwtErr, newJwt := GenerateToken(userEntity)
	if jwtErr != nil {
		return "", errors.New("error while creating new jwt")
	}

	return newJwt, nil
}

func GetUserIdFromToken(token string) (userId uint, err error) {
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil && err.Error() != "token has invalid claims: token is expired" {
		log.Error("Error while parsing token", err)
		return 0, err
	}

	userId = uint(claims["userId"].(float64))

	return userId, nil
}

func getJitFromToken(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil && err.Error() != "token has invalid claims: token is expired" {
		log.Error("Error while parsing token", err)
		return "", err
	}

	var jit = claims["jit"].(string)

	return jit, nil

}
