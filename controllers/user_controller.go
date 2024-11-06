package controllers

import (
	"golang-auth/repositories"
	"golang-auth/services"
	"golang-auth/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateNewRefreshToken(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Session not found"})
		return
	}
	storedRefresh, err := repositories.GetRefreshToken(sessionId)
	if err != nil || storedRefresh == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
	}

	claims, err := utils.ValidateToken(storedRefresh)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid session"})
	}
	userId := claims["user_id"].(string)
	role := claims["roles"].(string)
	newAccessToken, err := services.GenerateAccessToken(userId, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})

		return
	}
	refreshToken, err := services.GenerateRefreshToken(userId, role, time.Hour*10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	} else {
		repositories.UpdateRefreshToken(sessionId, refreshToken, time.Hour*10)
	}
	c.SetCookie("access_token", newAccessToken, int(time.Hour.Seconds()), "/", "localhost:8080", true, true)
	c.JSON(http.StatusOK, gin.H{"error": "sessionid created successfully"})

}
