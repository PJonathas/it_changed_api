package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Collection
var ctx context.Context

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_CONN_STR"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the database and collection variables
	db = client.Database("it_changed").Collection("changes")

	// TODO: Handle context
	ctx = context.TODO()
}
