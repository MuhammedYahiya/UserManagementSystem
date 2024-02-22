package routes

import (
	"github.com/MuhammedYahiya/UserManagementSystem/controllers"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.Singup)
		userRouter.POST("/login", controllers.LoginUser)
	}
}
