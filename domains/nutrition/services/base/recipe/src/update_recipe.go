package src

import (
	"context"
	"fmt"

	pbRecipe "github.com/kirvader/BodyController/domains/nutrition/services/base/recipe/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *RecipeService) UpdateRecipe(ctx context.Context, req *pbRecipe.UpdateRecipeRequest) (*pbRecipe.UpdateRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	newIngredientData, err := req.NewRecipeInfo.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	_, err = coll.UpdateByID(ctx, objectId, bson.M{"$set": newIngredientData})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	return &pbRecipe.UpdateRecipeResponse{}, nil
}
