package src

import (
	"context"

	pb "github.com/kirvader/BodyController/domains/nutrition/proto"
	ingredient "github.com/kirvader/BodyController/domains/nutrition/services/ingredient"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (svc *NutritionService) GetIngredient(ctx context.Context, req *pb.GetIngredientRequest) (*pb.GetIngredientResponse, error) {
	ingredientServiceResponse, err := svc.ingredientService.Get(ctx, &ingredient.GetIngredientRequest{
		HexId: req.GetIngredientHexId(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.GetIngredientResponse{
		Ingredient: ingredientServiceResponse.Ingredient,
	}, nil
}

func (svc *NutritionService) ListIngredients(ctx context.Context, req *pb.ListIngredientsRequest) (*pb.ListIngredientsResponse, error) {
	var lastPageToken *string = nil
	if req.LastPageToken != nil {
		lastPageToken = &req.LastPageToken.Value
	}

	ingredientServiceResponse, err := svc.ingredientService.List(ctx, &ingredient.ListIngredientsRequest{
		PageSize:      req.PageSize,
		LastPageToken: lastPageToken,
	})
	if err != nil {
		return nil, err
	}

	if ingredientServiceResponse.RetrievedPageToken == nil {
		return &pb.ListIngredientsResponse{
			Ingredients: ingredientServiceResponse.Ingredients,
		}, nil
	}

	return &pb.ListIngredientsResponse{
		Ingredients: ingredientServiceResponse.Ingredients,
		RetrievedPageToken: &wrapperspb.StringValue{
			Value: *ingredientServiceResponse.RetrievedPageToken,
		},
	}, nil
}
