package src

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kirvader/BodyController/internal/db"
)

type RecipeService struct {
	mongoClient *mongo.Client
}

func NewRecipeService(ctx context.Context) (*RecipeService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return nil, func() {}, err
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		return nil, func() {
			disconnectMongoClient()
		}, err
	}

	return &RecipeService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}
