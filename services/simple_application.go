package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func GetDatabase() (database *mongo.Database, ctx context.Context, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo_example:27017"))
	if err != nil {
		log.Println("database connection error", err)
		return nil, nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("err", err)
		return
	}
	log.Println("Successfully connected and pinged.")

	dbName := "mongo_example"
	database = client.Database(dbName)

	log.Println(dbName, database.Name())
	return
}

func main() {
	_, _, err := GetDatabase()
	if err != nil {
		return
	}
}
