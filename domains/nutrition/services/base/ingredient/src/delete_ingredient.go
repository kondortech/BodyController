package src

import (
	"context"
	"fmt"

	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (svc *IngredientService) DeleteIngredient(ctx context.Context, req *pbIngredient.DeleteIngredientRequest) (*pbIngredient.DeleteIngredientResponse, error) {
	ingredientsCollection := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	objectId, err := primitive.ObjectIDFromHex(req.HexIngredientId)
	if err != nil {
		return nil, err
	}

	_, err = ingredientsCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, fmt.Errorf("delete error: %w", err)
	}

	return &pbIngredient.DeleteIngredientResponse{}, nil
}
