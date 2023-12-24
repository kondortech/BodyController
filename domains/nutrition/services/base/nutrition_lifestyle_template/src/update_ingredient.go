package src

import (
	"context"
	"fmt"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/nutrition_lifestyle_template/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *NutritionLifestyleTemplateService) UpdateIngredient(ctx context.Context, req *pb.UpdateNutritionLifestyleTemplateRequest) (*pb.UpdateNutritionLifestyleTemplateResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("NutritionLifestyleTemplates")

	objectId, err := primitive.ObjectIDFromHex(req.NutritionLifestyleTemplateHexId)
	if err != nil {
		return nil, err
	}

	newDocumentData, err := req.NewNutritionLifestyleTemplateInfo.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	_, err = coll.UpdateByID(ctx, objectId, bson.M{"$set": newDocumentData})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	return &pb.UpdateNutritionLifestyleTemplateResponse{}, nil
}
