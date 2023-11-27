package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (svc *IngredientService) GetIngredient(ctx context.Context, req *pbIngredient.GetIngredientRequest) (*pbIngredient.GetIngredientResponse, error) {
	ingredientsCollection := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	objectId, err := primitive.ObjectIDFromHex(req.IngredientHexId)
	if err != nil {
		return nil, err
	}

	var ingredientMongoDB models.IngredientMongoDB
	err = ingredientsCollection.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		options.FindOne()).
		Decode(&ingredientMongoDB)
	if err != nil {
		return nil, err
	}

	ingredient, err := ingredientMongoDB.ConvertToProtoMessage()
	if err != nil {
		return nil, err
	}

	return &pbIngredient.GetIngredientResponse{
		Ingredient: ingredient,
	}, nil
}
