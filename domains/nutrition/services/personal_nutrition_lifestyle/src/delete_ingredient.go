package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/personal_nutrition_lifestyle/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *PersonalNutritionLifestyleService) DeleteIngredient(ctx context.Context, req *pb.DeletePersonalNutritionLifestyleRequest) (*pb.DeletePersonalNutritionLifestyleResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

	objectId, err := primitive.ObjectIDFromHex(req.PersonalNutritionLifestyleHexId)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &pb.DeletePersonalNutritionLifestyleResponse{}, nil
}
