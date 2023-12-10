package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pbRecipe "github.com/kirvader/BodyController/domains/nutrition/services/base/recipe/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (svc *RecipeService) GetRecipe(ctx context.Context, req *pbRecipe.GetRecipeRequest) (*pbRecipe.GetRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	var recipeMongoDB models.RecipeMongoDB
	err = coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		options.FindOne()).
		Decode(&recipeMongoDB)
	if err != nil {
		return nil, err
	}

	recipe, err := recipeMongoDB.ConvertToProtoMessage()
	if err != nil {
		return nil, err
	}

	return &pbRecipe.GetRecipeResponse{
		Recipe: recipe,
	}, nil
}
