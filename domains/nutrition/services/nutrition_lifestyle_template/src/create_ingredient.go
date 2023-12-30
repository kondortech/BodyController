package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/nutrition_lifestyle_template/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *NutritionLifestyleTemplateService) CreateIngredient(ctx context.Context, req *pb.CreateNutritionLifestyleTemplateRequest) (*pb.CreateNutritionLifestyleTemplateResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("NutritionLifestyleTemplates")

	mongoIngredient, err := req.NutritionLifestyleTemplate.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongoIngredient)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &pb.CreateNutritionLifestyleTemplateResponse{
		NutritionLifestyleTemplateHexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
