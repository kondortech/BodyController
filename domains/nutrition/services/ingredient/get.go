package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetIngredientRequest struct {
	HexId string
}

type GetIngredientResponse struct {
	Ingredient *models.Ingredient
}

func (svc *Service) Get(ctx context.Context, req *GetIngredientRequest) (*GetIngredientResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	var mongo models.IngredientMongoDB
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

	return &GetIngredientResponse{
		Ingredient: proto,
	}, nil
}
