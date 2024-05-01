package controllers

import (
	"dms/managers"
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func GenerateToken(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	parsingError := c.BindJSON(&body)
	if parsingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var dbError, userFromDb = repositories.GetByEmail(body.Email)
	if dbError != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email or password"})
		return
	}

	var password = bcrypt.CompareHashAndPassword([]byte(userFromDb.Password), []byte(body.Password))
	if password != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid email or password"})
		return
	}

	var err, token = managers.GenerateToken(userFromDb)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	var user = models.UserModel{
		ID:          userFromDb.ID,
		FirstName:   userFromDb.FirstName,
		LastName:    userFromDb.LastName,
		Email:       userFromDb.Email,
		PhoneNumber: userFromDb.PhoneNumber,
		DmsApiKey:   userFromDb.DmsApiKey,
		Country:     userFromDb.Country,
		BirthDate:   userFromDb.BirthDate.Format("2006-01-02"),
	}

	var refreshToken, refreshTokenError = managers.GenerateRefreshToken(token)
	if refreshTokenError != nil {
		log.Error("Error while generating refresh token", refreshTokenError)
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "refreshToken": refreshToken, "user": user})

}

func GetMe(c *gin.Context) {
	token := c.GetHeader("Authorization")

	var err, user = managers.GetUserFromValidToken(token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func RefreshToken(c *gin.Context) {

	var body struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refreshToken"`
	}

	parsingError := c.BindJSON(&body)

	if parsingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	newjwt, refreshError := managers.RefreshToken(body.Token, body.RefreshToken)

	if refreshError != nil {
		log.Errorf("Error while refreshing token, %v", refreshError)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid request"})
		return
	}
	var newRefreshToken, newRefreshTokenError = managers.GenerateRefreshToken(newjwt)
	if newRefreshTokenError != nil {
		log.Error("Error while generating new refresh token", newRefreshTokenError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	repositories.InvalidateTokenByValue(body.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"token": newjwt, "refreshToken": newRefreshToken})

}

func RegenerateApiKey(c *gin.Context) {
	err, user := managers.GetUserFromValidToken(c.GetHeader("Authorization"))
	if err != nil {
		return
	}
	newApiKey, err := services.RegenerateApiKeyForUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"apiKey": newApiKey})
}
