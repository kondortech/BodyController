package src

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeletePersonalNutritionLifestyleRequest struct {
	HexId string
}

type DeletePersonalNutritionLifestyleResponse struct{}

func (svc *PersonalNutritionLifestyleService) Delete(ctx context.Context, req *DeletePersonalNutritionLifestyleRequest) (*DeletePersonalNutritionLifestyleResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &DeletePersonalNutritionLifestyleResponse{}, nil
}
