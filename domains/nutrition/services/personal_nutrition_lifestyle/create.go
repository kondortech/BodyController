package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePersonalNutritionLifestyleRequest struct {
	PersonalNutritionLifestyle *models.PersonalNutritionLifestyle
}

type CreatePersonalNutritionLifestyleResponse struct {
	HexId string
}

func (svc *Service) Create(ctx context.Context, req *CreatePersonalNutritionLifestyleRequest) (*CreatePersonalNutritionLifestyleResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

	mongo, err := req.PersonalNutritionLifestyle.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongo)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &CreatePersonalNutritionLifestyleResponse{
		HexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
