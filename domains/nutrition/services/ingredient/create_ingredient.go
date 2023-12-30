package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateIngredientRequest struct {
	Ingredient *models.Ingredient
}

type CreateIngredientResponse struct {
	HexId string
}

func (svc *IngredientService) CreateIngredient(ctx context.Context, req *CreateIngredientRequest) (*CreateIngredientResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	mongo, err := req.Ingredient.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongo)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &CreateIngredientResponse{
		HexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
