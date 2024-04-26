package controllers

import (
	"dms/entites"
	"dms/managers"
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
)

func AddUser(c *gin.Context) {

	body := models.AddUserModel{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error while hashing password", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	birthDate, err := time.Parse("2006-01-02", body.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user := entites.User{
		FirstName:              body.FirstName,
		LastName:               body.LastName,
		Email:                  body.Email,
		Password:               string(password),
		PhoneNumber:            body.PhoneNumber,
		DmsApiKey:              services.GenerateApiKey(),
		Country:                body.Country,
		BirthDate:              birthDate,
		IsPhoneNumberConfirmed: false,
	}
	createdUser, err := repositories.AddUser(user)
	log.Infof("User created with id: %v", createdUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with provided email already exists"})
		return
	}

	var tokenErr, token = managers.GenerateToken(createdUser)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	var userModel = models.UserModel{
		ID:          createdUser.ID,
		FirstName:   createdUser.FirstName,
		LastName:    createdUser.LastName,
		Email:       createdUser.Email,
		PhoneNumber: createdUser.PhoneNumber,
		DmsApiKey:   createdUser.DmsApiKey,
	}

	c.JSON(http.StatusCreated, gin.H{"token": token, "apiKey": user.DmsApiKey, "user": userModel})

}

func GetUserById(c *gin.Context) {
	idStr := c.Param("userId")
	id, err := strconv.Atoi(idStr)
	user, userError := repositories.GetById(uint(id))
	if userError != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userModel := models.GetUserModel{
		Id:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Country:     user.Country,
		BirthDate:   user.BirthDate.Format("2006-01-02"),
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, userModel)
}

func UpdateUser(c *gin.Context) {
	var body models.AddUserModel
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	idStr := c.Param("userId")
	id, err := strconv.Atoi(idStr)

	user, userError := repositories.GetById(uint(id))
	if userError != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.PhoneNumber != body.PhoneNumber {
		user.IsPhoneNumberConfirmed = false
	}

	birthDate, err := time.Parse("2006-01-02", body.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName
	user.Email = body.Email
	user.PhoneNumber = body.PhoneNumber
	user.Country = body.Country
	user.BirthDate = birthDate

	_, err = repositories.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
