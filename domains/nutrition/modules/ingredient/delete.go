package src

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteIngredientRequest struct {
	HexId string
}

type DeleteIngredientResponse struct{}

func (svc *Service) Delete(ctx context.Context, req *DeleteIngredientRequest) (*DeleteIngredientResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	resp, err := coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	if resp.DeletedCount != 1 {
		return nil, errors.New("'delete ingredient' operation didn't succeed: unknown error see mongodb logs for more information")
	}

	return &DeleteIngredientResponse{}, nil
}
