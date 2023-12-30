package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateRecipeRequest struct {
	HexId   string
	NewData *models.Recipe
}

type UpdateRecipeResponse struct {
	HexId string
}

func (svc *Service) Update(ctx context.Context, req *UpdateRecipeRequest) (*UpdateRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	mongoNewInstance, err := req.NewData.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	resp, err := coll.UpdateByID(ctx, objectId, bson.M{"$set": mongoNewInstance})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}
	if resp.ModifiedCount == 0 {
		return nil, fmt.Errorf("no record updated: id='%s'", req.HexId)
	}

	return &UpdateRecipeResponse{
		HexId: req.HexId,
	}, nil
}
