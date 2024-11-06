package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user in the system.
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"user_id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"-"`
	Roles    string             `bson:"role" json:"role"`
}

// LoginRequest represents the structure for user login requests.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest represents the structure for user registration requests.
type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Roles    string `json:"role" binding:"required"` // Ensure roles are provided
}

type UserDetails struct {
	Id    string `json:"user_id"`
	Email string `json:"email"`
	Roles string `json:"role"`
}

type PasswordResetRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type PasswordReset struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
}
