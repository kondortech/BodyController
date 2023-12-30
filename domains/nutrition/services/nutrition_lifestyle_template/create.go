package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateNutritionLifestyleTemplateRequest struct {
	NutritionLifestyleTemplate *models.NutritionLifestyleTemplate
}

type CreateNutritionLifestyleTemplateResponse struct {
	HexId string
}

func (svc *Service) Create(ctx context.Context, req *CreateNutritionLifestyleTemplateRequest) (*CreateNutritionLifestyleTemplateResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("NutritionLifestyleTemplates")

	mongo, err := req.NutritionLifestyleTemplate.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongo)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &CreateNutritionLifestyleTemplateResponse{
		HexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
