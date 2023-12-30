package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"github.com/kirvader/BodyController/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListIngredientsRequest struct {
	PageSize      int32
	LastPageToken *string
}

type ListIngredientsResponse struct {
	Ingredients        []*models.Ingredient
	RetrievedPageToken *string
}

const DefaultPageSize = 10

func (svc *IngredientService) ListIngredients(ctx context.Context, req *ListIngredientsRequest) (*ListIngredientsResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = DefaultPageSize
	}

	coll := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	var currentPage *utils.Page
	if req.LastPageToken == nil {
		currentPage = &utils.Page{
			PageSize:   req.PageSize,
			PageOffset: 0,
		}
	} else {
		lastRetrievedPage, err := utils.PageFromToken(*req.LastPageToken)
		if err != nil {
			return nil, fmt.Errorf("error decoding last page token: %v", err)
		}
		currentPage = &utils.Page{
			PageSize:   req.PageSize,
			PageOffset: lastRetrievedPage.PageOffset + lastRetrievedPage.PageSize,
		}
	}

	options := options.Find()
	// TODO add filters, maybe also by taste - so waiting
	options.SetSort(bson.M{"macros_100g.calories": 1})
	options.SetSkip(int64(currentPage.PageOffset))
	options.SetLimit(int64(currentPage.PageSize))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	result := make([]*models.Ingredient, 0, currentPage.PageSize)

	for cursor.Next(ctx) {
		var mongoIngredient models.IngredientMongoDB
		err := cursor.Decode(&mongoIngredient)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %v", err)
		}
		ingredient, err := mongoIngredient.ConvertToProtoMessage()
		if err != nil {
			return nil, fmt.Errorf("error parsing mongo ingredient: %v", err)
		}
		result = append(result, ingredient)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	if len(result) < int(req.PageSize) {
		return &ListIngredientsResponse{
			Ingredients: result,
		}, nil
	}

	currentPageToken, err := currentPage.GetToken()
	if err != nil {
		return nil, fmt.Errorf("error when forming page token: %v", err)
	}

	return &ListIngredientsResponse{
		Ingredients:        result,
		RetrievedPageToken: &currentPageToken,
	}, nil
}
