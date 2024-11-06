package services

import (
	"golang-auth/models"
	"golang-auth/repositories"
	"golang-auth/utils"
)

func CreateUser(user *models.User) error {
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	return repositories.CreateUser(user)

}

func HashPassword(password string) (string, error) {
	return utils.GenerateHashPassword(password)
}
