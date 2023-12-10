package src

import (
	"context"
	"fmt"

	pbRecipe "github.com/kirvader/BodyController/domains/nutrition/services/base/recipe/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *RecipeService) DeleteRecipe(ctx context.Context, req *pbRecipe.DeleteRecipeRequest) (*pbRecipe.DeleteRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &pbRecipe.DeleteRecipeResponse{}, nil
}
