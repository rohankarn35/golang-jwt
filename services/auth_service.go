package services

import (
	"errors"
	"golang-auth/models"
	"golang-auth/repositories"
	"golang-auth/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthenticateUser(email, password string) (*models.User, error) {
	user, err := repositories.FindUserbyEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if !utils.ValidatePassword(user.Password, password) {
		return nil, errors.New("password not matched")
	}
	return user, nil

}

func GenerateAccessToken(userId, role string) (string, error) {
	return utils.GenerateJWT(userId, role, time.Hour*1)
}

func GenerateRefreshToken(userId, role string, expiration time.Duration) (string, error) {
	return utils.GenerateJWT(userId, role, expiration)
}

func RegisterUser(register *models.RegisterRequest) error {
	existingUser, err := repositories.FindUserbyEmail(register.Email)
	if err != nil {
		return err
	}

	if existingUser != nil {
		return errors.New("user already exist")
	}
	registeredUser := &models.User{
		ID:       primitive.NewObjectID(),
		Email:    register.Email,
		Password: register.Password,
		Roles:    register.Roles,
	}
	return CreateUser(registeredUser)

}

func RequestResetPassword(email string) error {
	user, err := repositories.FindUserbyEmail(email)
	if err != nil {
		return errors.New("email not found")
	}

	resetToken, err := utils.GenerateResetToken(user.ID.Hex())
	if err != nil {
		return errors.New("failed to generate reset token")
	}
	if err := utils.SendResetEmail(email, resetToken); err != nil {
		return errors.New("failed to send email")
	}

	return nil

}

func ResetPassword(token, newpassword string) error {
	claims, err := utils.ValidateToken(token)
	if err != nil {
		return errors.New("invalid or expired token")
	}
	userId, ok := claims["user_id"].(string)
	if !ok {
		return errors.New("invalid token")
	}
	hashedPassword, err := utils.GenerateHashPassword(newpassword)
	if err != nil {
		return errors.New("failed to generate hashpassword")
	}

	return repositories.UpdatePassword(userId, hashedPassword)

}
