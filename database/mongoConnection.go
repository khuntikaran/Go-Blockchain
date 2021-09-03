package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	MongoDb := os.Getenv("MONGODB_URL")
	clientOptions := options.Client().ApplyURI(MongoDb)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb.")

	collection := client.Database("chainDB").Collection("chain")
	return collection
}
