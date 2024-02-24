package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MuhammedYahiya/UserManagementSystem/models"
	"github.com/MuhammedYahiya/UserManagementSystem/utils"
	"github.com/dgrijalva/jwt-go"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": existingUser.ID,
		"email":  existingUser.Email,
	})

	tokenString, err := token.SignedString([]byte("https://jwt.io/#debugger-io?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6InVzZXIiLCJpYXQiOjE1MTYyMzkwMjJ9.khRbDuF1o5ZBSuM94UqI7sS-r6knwoHUDrI6-whE76E"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not loging"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}

func PostUpdateUser(c *gin.Context) {
	id := c.Param("id")
	type updateUserRequestBody struct {
		Username string `json:"username"`
		Email    string `json:"email" gorm:"unique"`
	}

	body := updateUserRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tokenUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not exists",
		})
		return
	}

	// Assert tokenUserID to float64
	tokenUserIDFloat, ok := tokenUserID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "not ok",
		})
		return
	}

	// Convert float64 to int
	tokenUserIDInt := int(tokenUserIDFloat)

	// Convert id to int for comparison
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error parsing id",
		})
		return
	}

	if tokenUserIDInt != idInt {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "id is incorrect",
		})
		return
	}

	var user models.User
	if result := utils.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if body.Email != "" {
		user.Email = body.Email
	}
	if body.Username != "" {
		user.Username = body.Username
	}

	utils.DB.Save(&user)
	c.JSON(http.StatusOK, &user)
}
