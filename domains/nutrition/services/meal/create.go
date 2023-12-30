package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateMealRequest struct {
	Meal *models.Meal
}

type CreateMealResponse struct {
	HexId string
}

func (svc *Service) Create(ctx context.Context, req *CreateMealRequest) (*CreateMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	mongo, err := req.Meal.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.InsertOne(ctx, mongo)
	if err != nil {
		return nil, fmt.Errorf("insert error: %w", err)
	}

	return &CreateMealResponse{
		HexId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}
