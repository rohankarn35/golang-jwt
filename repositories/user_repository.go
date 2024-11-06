package repositories

import (
	"context"
	"golang-auth/config"
	"golang-auth/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func Init() {
	userCollection = config.MongoDb.Database("go-auth").Collection("users")
}

func CreateUser(user *models.User) error {
	_, err := userCollection.InsertOne(config.Ctx, user)
	return err
}

func FindUserbyEmail(email string) (*models.User, error) {
	var user models.User
	err := userCollection.FindOne(config.Ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil

}

func FindUserId(userId string) (*models.User, error) {
	var user models.User
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	err = userCollection.FindOne(config.Ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil

}

func UpdatePassword(userId, hashedPassword string) error {
	filter := bson.M{"_id": userId}
	update := bson.M{"$set": bson.M{"password": hashedPassword}}
	_, err := userCollection.UpdateOne(context.TODO(), filter, update)
	return err
}
