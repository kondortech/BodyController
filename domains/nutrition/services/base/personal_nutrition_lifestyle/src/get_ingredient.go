package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/personal_nutrition_lifestyle/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (svc *PersonalNutritionLifestyleService) GetIngredient(ctx context.Context, req *pb.GetPersonalNutritionLifestyleRequest) (*pb.GetPersonalNutritionLifestyleResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

	objectId, err := primitive.ObjectIDFromHex(req.PersonalNutritionLifestyleHexId)
	if err != nil {
		return nil, err
	}

	var instance models.PersonalNutritionLifestyleMongoDB
	err = coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		options.FindOne()).
		Decode(&instance)
	if err != nil {
		return nil, err
	}

	result, err := instance.ConvertToProtoMessage()
	if err != nil {
		return nil, err
	}

	return &pb.GetPersonalNutritionLifestyleResponse{
		PersonalNutritionLifestyle: result,
	}, nil
}
