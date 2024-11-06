package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to generate password hash: %w", err)
	}

	return string(hashBytes), nil
}

func ValidatePassword(hashpassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}
