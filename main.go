package main

import (
	"github.com/MuhammedYahiya/UserManagementSystem/models"
	"github.com/MuhammedYahiya/UserManagementSystem/routes"
	"github.com/MuhammedYahiya/UserManagementSystem/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	utils.ConnectDB()
	err := utils.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	routes.InitializeRouter(r)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
