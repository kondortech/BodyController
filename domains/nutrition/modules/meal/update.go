package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateMealRequest struct {
	HexId   string
	NewData *models.Meal
}

type UpdateMealResponse struct {
	HexId string
}

func (svc *Service) UpdateIngredient(ctx context.Context, req *UpdateMealRequest) (*UpdateMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	newDocumentData, err := req.NewData.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.UpdateByID(ctx, objectId, bson.M{"$set": newDocumentData})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}
	if resp.ModifiedCount == 0 {
		return nil, fmt.Errorf("no record updated: id='%s'", req.HexId)
	}

	return &UpdateMealResponse{
		HexId: req.HexId,
	}, nil
}
