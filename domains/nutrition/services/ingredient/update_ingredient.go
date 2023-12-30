package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateIngredientRequest struct {
	HexId   string
	NewData *models.Ingredient
}

type UpdateIngredientResponse struct {
	HexId string
}

func (svc *IngredientService) UpdateIngredient(ctx context.Context, req *UpdateIngredientRequest) (*UpdateIngredientResponse, error) {
	ingredientsCollection := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	newIngredientData, err := req.NewData.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := ingredientsCollection.UpdateByID(ctx, objectId, bson.M{"$set": newIngredientData})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}
	if resp.ModifiedCount == 0 {
		return nil, fmt.Errorf("no record updated: id='%s'", req.HexId)
	}

	return &UpdateIngredientResponse{
		HexId: req.HexId,
	}, nil
}
