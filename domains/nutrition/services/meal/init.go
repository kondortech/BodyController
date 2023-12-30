package src

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kirvader/BodyController/internal/db"
)

type MealService struct {
	mongoClient *mongo.Client
}

func NewMealService(ctx context.Context) (*MealService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return nil, func() {}, err
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		return nil, func() {
			disconnectMongoClient()
		}, err
	}

	return &MealService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}
