package middleware

import (
	"golang-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Auth token not provided"})
			ctx.Abort()

			return
		}
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}
		claims, err := utils.ValidateToken(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}
		ctx.Set("userID", claims["user_id"])
		ctx.Set("role", claims["role"])

		ctx.Next()

	}
}
