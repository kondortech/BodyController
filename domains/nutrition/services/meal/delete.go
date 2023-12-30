package src

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteMealRequest struct {
	HexId string
}

type DeleteMealResponse struct{}

func (svc *MealService) Delete(ctx context.Context, req *DeleteMealRequest) (*DeleteMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &DeleteMealResponse{}, nil
}
