package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateRecipeRequest struct {
	Recipe *models.Recipe
}

type CreateRecipeResponse struct {
	HexId string
}

func (svc *RecipeService) Create(ctx context.Context, req *CreateRecipeRequest) (*CreateRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	mongo, err := req.Recipe.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongo)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &CreateRecipeResponse{
		HexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
