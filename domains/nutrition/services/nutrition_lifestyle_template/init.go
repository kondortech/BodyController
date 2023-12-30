package src

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kirvader/BodyController/internal/db"
)

type NutritionLifestyleTemplateService struct {
	mongoClient *mongo.Client
}

func NewNutritionLifestyleTemplateService(ctx context.Context) (*NutritionLifestyleTemplateService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return nil, func() {}, err
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		return nil, func() {
			disconnectMongoClient()
		}, err
	}

	return &NutritionLifestyleTemplateService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}
