package main

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

func main() {
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017", nil)

	if err != nil {
		log.Fatalf("Error opening connectio to MongoDB: %v", err)
	}

	db := client.Database("go-microservice-httlp")

	
}
