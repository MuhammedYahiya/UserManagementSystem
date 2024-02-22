package controllers

import (
	"github.com/MuhammedYahiya/UserManagementSystem/models"
	"github.com/MuhammedYahiya/UserManagementSystem/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Singup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashPassword)
	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"message": "User registered successfully"})
}
