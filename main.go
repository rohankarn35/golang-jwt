package main

import (
	"golang-auth/config"
	"golang-auth/repositories"
	"golang-auth/routes"
	"golang-auth/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()

	utils.Init()

	repositories.Init()

	router := gin.Default()

	routes.SetupRoutes(router)
	port := "8080"
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
