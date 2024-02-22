package controllers

import (
	"fmt"
	"net/http"

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

func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Printf("User: %+v\n", user)
	var existingUser models.User
	if err := utils.DB.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username"})
		return
	}

	// fmt.Println(existingUser.Password)
	fmt.Println(user.Password)
	// Compare the hashed password with the plain text password
	err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		// Password mismatch, return appropriate error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
