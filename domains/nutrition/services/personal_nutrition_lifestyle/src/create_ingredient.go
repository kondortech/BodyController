package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/personal_nutrition_lifestyle/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *PersonalNutritionLifestyleService) CreateIngredient(ctx context.Context, req *pb.CreatePersonalNutritionLifestyleRequest) (*pb.CreatePersonalNutritionLifestyleResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

	mongoPNL, err := req.PersonalNutritionLifestyle.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongoPNL)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &pb.CreatePersonalNutritionLifestyleResponse{
		PersonalNutritionLifestyleHexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
