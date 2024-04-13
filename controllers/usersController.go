package controllers

import (
	"dms/entites"
	"dms/models"
	"dms/repositories"
	"dms/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func AddUser(c *gin.Context) {

	userModel := models.AddUserModel{}
	err := c.BindJSON(&userModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	password, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error while hashing password", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Please try again later"})
		return
	}

	user := entites.User{
		FirstName:   userModel.FirstName,
		LastName:    userModel.LastName,
		Email:       userModel.Email,
		Password:    string(password),
		PhoneNumber: userModel.PhoneNumber,
		DmsApiKey:   services.GenerateApiKey(),
	}
	userId, err := repositories.AddUser(user)
	log.Infof("User created with id: %v", userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with provided email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUserById(c *gin.Context) {
	idStr := c.Param("userId")
	id, err := strconv.Atoi(idStr)
	user, err := repositories.GetById(id)

	userModel := models.GetUserModel{}

	userModel.FirstName = user.FirstName
	userModel.LastName = user.LastName
	userModel.Email = user.Email
	userModel.PhoneNumber = user.PhoneNumber
	userModel.CreatedAt = user.CreatedAt

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
