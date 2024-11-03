package config

import (
	"fmt"
	"golang-auth/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	verifiedRoles = utils.VerifiedRoles
	jwtSecret     string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	jwtSecret = os.Getenv("JWTSECRET")
	if jwtSecret == "" {
		panic("JWT_SECRET cannot be empty")
	}
}

func isVerifiedRole(role string) bool {
	for _, r := range verifiedRoles {
		if r == role {
			return true
		}
	}
	return false
}

func isValidUserId(userId string) bool {
	_, err := primitive.ObjectIDFromHex(userId)
	return err == nil
}

func GenerateJWT(userId string, role string, expiration time.Duration) (string, error) {
	if userId == "" {
		return "", fmt.Errorf("UserId cannot be empty")
	}
	if !isValidUserId(userId) {
		return "", fmt.Errorf("invalid userId")
	}
	if role == "" {
		return "", fmt.Errorf("role cannot be empty")
	}
	if !isVerifiedRole(role) {
		return "", fmt.Errorf("invalid role")
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"roles":   role,
		"exp":     time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error generating jwt: %v", err)
	}

	return signedToken, nil
}

func ValidateToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}
