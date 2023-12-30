package src

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteRecipeRequest struct {
	HexId string
}

type DeleteRecipeResponse struct{}

func (svc *RecipeService) DeleteRecipe(ctx context.Context, req *DeleteRecipeRequest) (*DeleteRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &DeleteRecipeResponse{}, nil
}
