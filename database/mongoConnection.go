package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	/*err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	MongoDb := os.Getenv("MONGODB_URL")*/
	clientOptions := options.Client().ApplyURI("mongodb+srv://karan:maher7505@cluster0.crnpa.mongodb.net/chainDB?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected to mongodb.")

	collection := client.Database("chainDB").Collection("chain")
	return collection
}
