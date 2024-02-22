package main

import (
	"github.com/MuhammedYahiya/UserManagementSystem/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	utils.ConnectDB()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "created server",
		})
	})

	r.Run()
}
