package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/meal/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *MealService) CreateIngredient(ctx context.Context, req *pb.CreateMealRequest) (*pb.CreateMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	mongoDocument, err := req.Meal.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongoDocument)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &pb.CreateMealResponse{
		MealHexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
