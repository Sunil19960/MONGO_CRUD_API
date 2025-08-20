package database

import (
	"context"
	"crud/utils"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(utils.Settings.Mongo_URI).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	utils.MongoClient, err = mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	_, err = PingMongoDB()
	if err != nil {
		log.Fatalf("Unable to connect to MongoDB Client")
	}
}

func PingMongoDB() (string, error) {
	deadline := time.Now().Add(10 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	err := utils.MongoClient.Ping(ctx, nil)
	if err != nil {
		return "Failed", err
	}
	return "Success", nil

}
