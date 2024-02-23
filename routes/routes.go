package routes

import (
	"github.com/MuhammedYahiya/UserManagementSystem/controllers"
	"github.com/MuhammedYahiya/UserManagementSystem/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/register", controllers.Singup)
		userRouter.POST("/login", controllers.LoginUser)
	}

	adminRouter := r.Group("/admin")
	{
		adminRouter.POST("/login", controllers.AdminLoginUser)
		adminRouter.GET("/users", middleware.AdminAuthMiddleware(), controllers.GetAllUsers)
		adminRouter.GET("/user/:id", middleware.AdminAuthMiddleware(), controllers.GetUserById)
	}
}
