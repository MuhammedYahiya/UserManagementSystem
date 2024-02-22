package controllers

import (
	"fmt"
	"net/http"

	"github.com/MuhammedYahiya/UserManagementSystem/models"
	"github.com/MuhammedYahiya/UserManagementSystem/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AdminLoginUser(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	fmt.Printf("User: %+v\n", admin)
	var existingAdmin models.Admin
	if err := utils.DB.Where("email = ?", admin.Email).First(&existingAdmin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// fmt.Println(existingUser.Password)
	fmt.Println(admin.Password)
	// Compare the hashed password with the plain text password
	err := bcrypt.CompareHashAndPassword([]byte(existingAdmin.Password), []byte(admin.Password))
	if err != nil {
		// Password mismatch, return appropriate error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
