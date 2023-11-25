package db

import (
	"context"
	"fmt"
	"log"

	"github.com/kirvader/BodyController/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InitMongoDBClientFromENV(ctx context.Context) (mongoClient *mongo.Client, closeFunction func(), err error) {
	mongoDBIP := utils.GetEnvWithDefault("MONGODB_IP", "0.0.0.0")
	mongoDBPort := utils.GetEnvWithDefault("MONGODB_PORT", "27017")

	log.Println("IP: ", mongoDBIP)
	log.Println("PORT: ", mongoDBPort)

	mongoDBURI := fmt.Sprintf("mongodb://%s:%s", mongoDBIP, mongoDBPort)

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoDBURI))
	if err != nil {
		return nil, func() {}, err
	}
	return mongoClient, func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}, nil
}

func PingMongoDb(ctx context.Context, mongoClient *mongo.Client) error {
	return mongoClient.Ping(ctx, readpref.Primary())
}
