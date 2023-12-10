package src

import (
	"context"
	"fmt"

	pbRecipe "github.com/kirvader/BodyController/domains/nutrition/services/base/recipe/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *RecipeService) CreateRecipe(ctx context.Context, req *pbRecipe.CreateRecipeRequest) (*pbRecipe.CreateRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	mongoIngredient, err := req.Recipe.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongoIngredient)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &pbRecipe.CreateRecipeResponse{
		HexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
