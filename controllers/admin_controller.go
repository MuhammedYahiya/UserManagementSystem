package controllers

import (
	"fmt"
	"net/http"

	"github.com/MuhammedYahiya/UserManagementSystem/models"
	"github.com/MuhammedYahiya/UserManagementSystem/utils"
	"github.com/dgrijalva/jwt-go"
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

	// Compare the hashed password with the plain text password
	err := bcrypt.CompareHashAndPassword([]byte(existingAdmin.Password), []byte(admin.Password))
	if err != nil {
		// Password mismatch, return appropriate error
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminID": existingAdmin.ID,
		"email":   existingAdmin.Email,
	})

	tokenString, err := token.SignedString([]byte("https://jwt.io/#debugger-io?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not loging"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
