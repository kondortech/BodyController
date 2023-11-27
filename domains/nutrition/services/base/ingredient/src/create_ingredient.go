package src

import (
	"context"
	"fmt"

	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *IngredientService) CreateIngredient(ctx context.Context, req *pbIngredient.CreateIngredientRequest) (*pbIngredient.CreateIngredientResponse, error) {
	ingredientsCollection := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	mongoIngredient, err := req.Ingredient.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := ingredientsCollection.InsertOne(ctx, mongoIngredient)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &pbIngredient.CreateIngredientResponse{
		HexIngredientId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
