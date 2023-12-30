package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"

	"github.com/kirvader/BodyController/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ListRecipesRequest struct {
	PageSize      int32
	LastPageToken *string
}

type ListRecipesResponse struct {
	Recipes            []*models.Recipe
	RetrievedPageToken *string
}

const DefaultPageSize = 10

func (svc *RecipeService) List(ctx context.Context, req *ListRecipesRequest) (*ListRecipesResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = DefaultPageSize
	}

	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	var pageToRetrieve *utils.Page
	if req.LastPageToken == nil {
		pageToRetrieve = &utils.Page{
			PageSize:   req.PageSize,
			PageOffset: 0,
		}
	} else {
		lastRetrievedPage, err := utils.PageFromToken(*req.LastPageToken)
		if err != nil {
			return nil, fmt.Errorf("error decoding last page token: %v", err)
		}
		pageToRetrieve = &utils.Page{
			PageSize:   req.PageSize,
			PageOffset: lastRetrievedPage.PageOffset + lastRetrievedPage.PageSize,
		}
	}

	options := options.Find()
	// TODO add filters, maybe also by taste - so waiting
	options.SetSort(bson.M{"title": 1})
	options.SetSkip(int64(pageToRetrieve.PageOffset))
	options.SetLimit(int64(pageToRetrieve.PageSize))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	result := make([]*models.Recipe, 0, pageToRetrieve.PageSize)

	for cursor.Next(ctx) {
		var mongoRecipe models.RecipeMongoDB
		err := cursor.Decode(&mongoRecipe)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %v", err)
		}
		recipe, err := mongoRecipe.ConvertToProtoMessage()
		if err != nil {
			return nil, fmt.Errorf("error parsing mongo recipe: %v", err)
		}
		result = append(result, recipe)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	if len(result) < int(req.PageSize) {
		return &ListRecipesResponse{
			Recipes: result,
		}, nil
	}

	currentPageToken, err := pageToRetrieve.GetToken()
	if err != nil {
		return nil, fmt.Errorf("error when forming page token: %v", err)
	}

	return &ListRecipesResponse{
		Recipes:            result,
		RetrievedPageToken: &currentPageToken,
	}, nil
}
