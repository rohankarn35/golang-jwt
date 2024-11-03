package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	UserID    primitive.ObjectID `bson:"user_id"`
	Token     string             `json:"token" bson:"token"`
	ExpiresAt time.Time          `json:"expires_at" bson:"expires_at"`
}

type JWTClaims struct {
	UserId     string        `json:"user_id"`
	Roles      string        `json:"roles"`
	Expiration time.Duration `json:"exp"`
}
