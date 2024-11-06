package controllers

import (
	"fmt"
	"golang-auth/models"
	"golang-auth/repositories"
	"golang-auth/services"
	"golang-auth/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginUser models.LoginRequest
	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	user, err := services.AuthenticateUser(loginUser.Email, loginUser.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	accessToken, err := services.GenerateAccessToken(user.ID.Hex(), user.Roles)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate access token"})
		return
	}

	sessionId, err := utils.GenerateSessionId()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate sessionId"})
		return
	}

	refreshToken, err := services.GenerateRefreshToken(user.ID.Hex(), user.Roles, time.Hour*10)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate refresh token"})
		return
	}

	if err := repositories.StoreRefreshToken(sessionId, refreshToken, time.Hour*10); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to store refresh token"})
		return

	}

	c.SetCookie("access_token", accessToken, int(time.Hour.Seconds()), "/", "localhost", true, true)
	c.SetCookie("session_id", sessionId, int(time.Hour.Seconds()), "/", "localhost", true, true)

	// Set role-based access cookies

	c.JSON(http.StatusOK, gin.H{"message": "Loggined Successfully"})

}

func Logout(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "sessionId not found"})
		return
	}

	storedRefresh, err := repositories.GetRefreshToken(sessionId)
	if err != nil || storedRefresh == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired session"})
	}

	if err := repositories.DeleteRefreshToken(sessionId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove refreshtoken"})
		return
	}

	c.SetCookie("access_token", "", -1, "/", "localhost:8080", true, true)
	c.SetCookie("session_id", "", -1, "/", "localhost:8080", true, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out Successfully"})

}

func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request data"})
		fmt.Println(err)
		return
	}
	err := services.RegisterUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})

}

func ForgotPassword(c *gin.Context) {
	var req models.PasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := services.RequestResetPassword(req.Email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reset email sent"})

}

func ResetPassword(c *gin.Context) {
	var req models.PasswordReset
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := services.ResetPassword(req.Token, req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "password reset successfully"})

}
