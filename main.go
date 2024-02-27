package main

import (
	"auth/controllers"
	"auth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvInitializers()
	initializers.DatabaseInitialize()
}

func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.Run()
}
