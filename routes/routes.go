package routes

import (
	"golang-auth/controllers"
	"golang-auth/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
	router.POST("/forgot-password", controllers.ForgotPassword)

	authRoutes := router.Group("/auth")
	authRoutes.Use(middleware.AuthMiddleWare())
	{
		authRoutes.POST("/logout", controllers.Logout)
		authRoutes.POST("/reset-password", controllers.ResetPassword)
		authRoutes.POST("/refresh", controllers.GenerateNewRefreshToken)
	}

}
