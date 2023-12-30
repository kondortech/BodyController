package src

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kirvader/BodyController/internal/db"
)

type Service struct {
	mongoClient *mongo.Client
}

func NewService(ctx context.Context) (*Service, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return nil, func() {}, err
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		return nil, func() {
			disconnectMongoClient()
		}, err
	}

	return &Service{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}
