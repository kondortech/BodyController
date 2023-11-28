package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
	"github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/src/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const DefaultPageSize = 10

func (svc *IngredientService) ListIngredients(ctx context.Context, req *pbIngredient.ListIngredientsRequest) (*pbIngredient.ListIngredientsResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = DefaultPageSize
	}

	ingredientsCollection := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	var pageToRetrieve *utils.Page
	if req.GetLastPageToken() == nil {
		pageToRetrieve = &utils.Page{
			PageSize:   req.GetPageSize(),
			PageOffset: 0,
		}
	} else {
		lastRetrievedPage, err := utils.PageFromToken(req.GetLastPageToken().GetValue())
		if err != nil {
			return nil, fmt.Errorf("error decoding last page token: %v", err)
		}
		pageToRetrieve = &utils.Page{
			PageSize:   req.GetPageSize(),
			PageOffset: lastRetrievedPage.PageOffset + lastRetrievedPage.PageSize,
		}
	}

	options := options.Find()
	// TODO add filters, maybe also by taste - so waiting
	options.SetSort(bson.M{"macros_100g.calories": 1})
	options.SetSkip(int64(pageToRetrieve.PageOffset))
	options.SetLimit(int64(pageToRetrieve.PageSize))

	cursor, err := ingredientsCollection.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	result := make([]*models.Ingredient, 0, pageToRetrieve.PageSize)

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
		return &pbIngredient.ListIngredientsResponse{
			Ingredients: result,
		}, nil
	}

	currentPageToken, err := pageToRetrieve.GetToken()
	if err != nil {
		return nil, fmt.Errorf("error when forming page token: %v", err)
	}

	return &pbIngredient.ListIngredientsResponse{
		Ingredients:        result,
		RetrievedPageToken: wrapperspb.String(currentPageToken),
	}, nil
}
