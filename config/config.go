package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Ctx     = context.Background()
	MongoDb *mongo.Client
	RedisDb *redis.Client
)

func InitConfig() {

	initMongoDb()
	initRedis()

}

func initMongoDb() {
	mongoUri := os.Getenv("MONGODB_URI")
	fmt.Print("loaded mongouri ", mongoUri)
	clientOptions := options.Client().ApplyURI(mongoUri)
	var err error
	MongoDb, err = mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}
	err = MongoDb.Ping(Ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}
	fmt.Println("Connected to MongoDB")
}

func initRedis() {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	fmt.Print("connection error")

	if err != nil {
		log.Fatal("Invalid REDIS_DB value:", err)
	}
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	_, err = RedisDb.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("Redis connection error:", err)
	}
	fmt.Println("Connected to Redis")
}
