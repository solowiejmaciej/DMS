package controllers

import (
	"dms/managers"
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})

}

func GetMe(c *gin.Context) {
	var body struct {
		Token string `json:"token"`
	}

	parsingError := c.BindJSON(&body)
	if parsingError != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var err, user = managers.GetUserFromToken(body.Token)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func RefreshToken(c *gin.Context) {

}

func RegenerateApiKey(c *gin.Context) {
	err, user := managers.GetUserFromToken(c.GetHeader("Authorization"))
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
