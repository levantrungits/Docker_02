package main

import (
	"fmt"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	fmt.Println("=====Lam` Quen Docker with GO.=====")
	ConnectMongoDb()
}

func ConnectMongoDb() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://host.docker.internal:27017"))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDb Connection Ok.")
}