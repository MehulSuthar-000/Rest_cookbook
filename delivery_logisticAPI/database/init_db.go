package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initialize mongo db connection

var Client *mongo.Client
var Database *mongo.Database

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Panicln("❌ Failed to connect to MongoDB:", err)
	}

	// ✅ Ping MongoDB
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Panicln("❌ MongoDB connection test failed:", err)
	} else {
		fmt.Println("🚀 Connected to MongoDB!")
	}

	// ✅ Assign global client & database
	Client = client
	Database = client.Database("logistics_db") // Change to your DB name
}

// GetCollection returns the collection from the database
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("LogisticAPI").Collection(collectionName)
}
