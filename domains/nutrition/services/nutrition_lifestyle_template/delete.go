package src

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteNutritionLifestyleTemplateRequest struct {
	HexId string
}

type DeleteNutritionLifestyleTemplateResponse struct{}

func (svc *Service) Delete(ctx context.Context, req *DeleteNutritionLifestyleTemplateRequest) (*DeleteNutritionLifestyleTemplateResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("NutritionLifestyleTemplates")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &DeleteNutritionLifestyleTemplateResponse{}, nil
}
