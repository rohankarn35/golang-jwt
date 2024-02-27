package main

import (
	"auth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvInitializers()
	initializers.DatabaseInitialize()
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"MESSAGE": "SUCCESSF",
		})
	})
	r.Run()
}
