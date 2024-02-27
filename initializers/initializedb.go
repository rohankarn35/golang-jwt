package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func DatabaseInitialize() {
	var err error
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Error: mongodb uri not found in .env file")
	}

	clientoptions := options.Client().ApplyURI(uri)
	Client, err = mongo.Connect(context.TODO(), clientoptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error verifying MongoDB connection:", err)
	}

	fmt.Println("MongoDB connection successful")

}
