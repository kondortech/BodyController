package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetMealRequest struct {
	HexId string
}

type GetMealResponse struct {
	Meal *models.Meal
}

func (svc *Service) Get(ctx context.Context, req *GetMealRequest) (*GetMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	var mongo models.MealMongo
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

	return &GetMealResponse{
		Meal: proto,
	}, nil
}
