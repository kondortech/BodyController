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
		panic(err)
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	return &RecipeService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}
