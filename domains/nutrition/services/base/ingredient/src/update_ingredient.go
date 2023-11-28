package src

import (
	"context"
	"fmt"

	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *IngredientService) UpdateIngredient(ctx context.Context, req *pbIngredient.UpdateIngredientRequest) (*pbIngredient.UpdateIngredientResponse, error) {
	ingredientsCollection := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	objectId, err := primitive.ObjectIDFromHex(req.IngredientHexId)
	if err != nil {
		return nil, err
	}

	newIngredientData, err := req.NewIngredientInfo.ConvertToMongoDocument()
	if err != nil {
		return nil, err
	}

	_, err = ingredientsCollection.UpdateByID(ctx, objectId, bson.M{"$set": newIngredientData})
	if err != nil {
		return nil, fmt.Errorf("update error: %w", err)
	}

	return &pbIngredient.UpdateIngredientResponse{}, nil
}
