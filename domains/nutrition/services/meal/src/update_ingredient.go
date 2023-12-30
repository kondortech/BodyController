package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/meal/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *MealService) UpdateIngredient(ctx context.Context, req *pb.UpdateMealRequest) (*pb.UpdateMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	objectId, err := primitive.ObjectIDFromHex(req.MealHexId)
	if err != nil {
		return nil, err
	}

	newDocumentData, err := req.NewMealInfo.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	_, err = coll.UpdateByID(ctx, objectId, bson.M{"$set": newDocumentData})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	return &pb.UpdateMealResponse{}, nil
}
