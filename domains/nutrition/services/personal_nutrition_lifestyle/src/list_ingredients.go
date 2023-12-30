package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/personal_nutrition_lifestyle/proto"
	"github.com/kirvader/BodyController/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const DefaultPageSize = 10

func (svc *PersonalNutritionLifestyleService) ListIngredients(ctx context.Context, req *pb.ListPersonalNutritionLifestylesRequest) (*pb.ListPersonalNutritionLifestylesResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = DefaultPageSize
	}

	coll := svc.mongoClient.Database("BodyController").Collection("PersonalNutritionLifestyles")

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

	result := make([]*models.PersonalNutritionLifestyle, 0, pageToRetrieve.PageSize)

	for cursor.Next(ctx) {
		var mongoInstance models.PersonalNutritionLifestyleMongoDB
		err := cursor.Decode(&mongoInstance)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %v", err)
		}
		protoInstance, err := mongoInstance.ConvertToProtoMessage()
		if err != nil {
			return nil, fmt.Errorf("error parsing mongo PersonalNutritionLifestyle: %v", err)
		}
		result = append(result, protoInstance)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	if len(result) < int(req.PageSize) {
		return &pb.ListPersonalNutritionLifestylesResponse{
			PersonalNutritionLifestyles: result,
		}, nil
	}

	currentPageToken, err := pageToRetrieve.GetToken()
	if err != nil {
		return nil, fmt.Errorf("error when forming page token: %v", err)
	}

	return &pb.ListPersonalNutritionLifestylesResponse{
		PersonalNutritionLifestyles: result,
		RetrievedPageToken:          wrapperspb.String(currentPageToken),
	}, nil
}
