package src

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kirvader/BodyController/internal/db"
)

type PersonalNutritionLifestyleService struct {
	mongoClient *mongo.Client
}

func NewPersonalNutritionLifestyleService(ctx context.Context) (*PersonalNutritionLifestyleService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return nil, func() {}, err
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		return nil, func() {
			disconnectMongoClient()
		}, err
	}

	return &PersonalNutritionLifestyleService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}
