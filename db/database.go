package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Collection *mongo.Collection

func init() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	Collection = Client.Database("library").Collection("items")
}
