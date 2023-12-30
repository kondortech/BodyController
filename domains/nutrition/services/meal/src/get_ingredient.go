package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/meal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (svc *MealService) GetIngredient(ctx context.Context, req *pb.GetMealRequest) (*pb.GetMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	objectId, err := primitive.ObjectIDFromHex(req.MealHexId)
	if err != nil {
		return nil, err
	}

	var mongoInstance models.MealMongo
	err = coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		options.FindOne()).
		Decode(&mongoInstance)
	if err != nil {
		return nil, err
	}

	result, err := mongoInstance.ConvertToProtoMessage()
	if err != nil {
		return nil, err
	}

	return &pb.GetMealResponse{
		Meal: result,
	}, nil
}
