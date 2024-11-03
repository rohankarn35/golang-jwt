package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	FullName string             `json:"fullname" bson:"fullname"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Role     string             `json:"role" bson:"role"`
}

type LoginUser struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
