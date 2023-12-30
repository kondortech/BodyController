package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetRecipeRequest struct {
	HexId string
}

type GetRecipeResponse struct {
	Recipe *models.Recipe
}

func (svc *RecipeService) Get(ctx context.Context, req *GetRecipeRequest) (*GetRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	var mongo models.RecipeMongoDB
	err = coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		options.FindOne()).
		Decode(&mongo)
	if err != nil {
		return nil, err
	}

	proto, err := mongo.ConvertToProtoMessage()
	if err != nil {
		return nil, err
	}

	return &GetRecipeResponse{
		Recipe: proto,
	}, nil
}
