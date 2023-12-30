package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/meal/proto"
	"github.com/kirvader/BodyController/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const DefaultPageSize = 10

func (svc *MealService) ListIngredients(ctx context.Context, req *pb.ListMealsRequest) (*pb.ListMealsResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = DefaultPageSize
	}

	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

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
	options.SetSort(bson.M{"username": 1})
	options.SetSkip(int64(pageToRetrieve.PageOffset))
	options.SetLimit(int64(pageToRetrieve.PageSize))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	result := make([]*models.Meal, 0, pageToRetrieve.PageSize)

	for cursor.Next(ctx) {
		var mongoInstance models.MealMongo
		err := cursor.Decode(&mongoInstance)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %v", err)
		}
		protoInstance, err := mongoInstance.ConvertToProtoMessage()
		if err != nil {
			return nil, fmt.Errorf("error parsing mongo Meal: %v", err)
		}
		result = append(result, protoInstance)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	if len(result) < int(req.PageSize) {
		return &pb.ListMealsResponse{
			Meals: result,
		}, nil
	}

	currentPageToken, err := pageToRetrieve.GetToken()
	if err != nil {
		return nil, fmt.Errorf("error when forming page token: %v", err)
	}

	return &pb.ListMealsResponse{
		Meals:              result,
		RetrievedPageToken: wrapperspb.String(currentPageToken),
	}, nil
}
