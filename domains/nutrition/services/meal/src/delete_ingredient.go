package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/meal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *MealService) DeleteIngredient(ctx context.Context, req *pb.DeleteMealRequest) (*pb.DeleteMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	objectId, err := primitive.ObjectIDFromHex(req.MealHexId)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &pb.DeleteMealResponse{}, nil
}
