package repository

import (
	"blockchain/database"
	"blockchain/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.ConnectDB()

type document struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Block models.Block       `bson:"block,omitempty"`
}

func AddBlock(block models.Block) {
	data := document{
		Block: block,
	}
	_, err := collection.InsertOne(context.TODO(), data)

	if err != nil {
		fmt.Println("error in inserting document in mongodb", err)
	}
	fmt.Println("Block stored successfully")
}
