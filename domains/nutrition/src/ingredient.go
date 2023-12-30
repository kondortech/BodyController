package src

import (
	"context"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/aggregation/proto"
	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
)

func (svc *NutritionService) GetIngredient(ctx context.Context, req *pb.GetIngredientRequest) (*pb.GetIngredientResponse, error) {
	ingredientServiceResponse, err := svc.ingredientServiceClient.GetIngredient(ctx, &pbIngredient.GetIngredientRequest{
		IngredientHexId: req.GetIngredientHexId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetIngredientResponse{
		Ingredient: ingredientServiceResponse.GetIngredient(),
	}, nil
}

func (svc *NutritionService) ListIngredients(ctx context.Context, req *pb.ListIngredientsRequest) (*pb.ListIngredientsResponse, error) {
	ingredientServiceResponse, err := svc.ingredientServiceClient.ListIngredients(ctx, &pbIngredient.ListIngredientsRequest{
		PageSize:      req.PageSize,
		LastPageToken: req.LastPageToken,
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListIngredientsResponse{
		Ingredients:        ingredientServiceResponse.Ingredients,
		RetrievedPageToken: ingredientServiceResponse.RetrievedPageToken,
	}, nil
}
