package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/personal_nutrition_lifestyle/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *PersonalNutritionLifestyleService) UpdateIngredient(ctx context.Context, req *pb.UpdatePersonalNutritionLifestyleRequest) (*pb.UpdatePersonalNutritionLifestyleResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

	objectId, err := primitive.ObjectIDFromHex(req.PersonalNutritionLifestyleHexId)
	if err != nil {
		return nil, err
	}

	newDocumentData, err := req.NewPersonalNutritionLifestyleInfo.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	_, err = coll.UpdateByID(ctx, objectId, bson.M{"$set": newDocumentData})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	return &pb.UpdatePersonalNutritionLifestyleResponse{}, nil
}
