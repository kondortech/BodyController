package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetPersonalNutritionLifestyleRequest struct {
	HexId string
}

type GetPersonalNutritionLifestyleResponse struct {
	PersonalNutritionLifestyle *models.PersonalNutritionLifestyle
}

func (svc *PersonalNutritionLifestyleService) GetIngredient(ctx context.Context, req *GetPersonalNutritionLifestyleRequest) (*GetPersonalNutritionLifestyleResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	var mongo models.PersonalNutritionLifestyleMongoDB
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

	return &GetPersonalNutritionLifestyleResponse{
		PersonalNutritionLifestyle: proto,
	}, nil
}
