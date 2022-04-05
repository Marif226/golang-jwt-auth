package database

import (
	"context"
	"fmt"
	"line"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/vuln/client"
)

func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongo.NewClient(options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}

	context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
	return collection
}