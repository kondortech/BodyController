package src

import (
	"context"
	"fmt"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/nutrition_lifestyle_template/proto"
	"github.com/kirvader/BodyController/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const DefaultPageSize = 10

func (svc *NutritionLifestyleTemplateService) ListIngredients(ctx context.Context, req *pb.ListNutritionLifestyleTemplatesRequest) (*pb.ListNutritionLifestyleTemplatesResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = DefaultPageSize
	}

	coll := svc.mongoClient.Database("BodyController").Collection("NutritionLifestyleTemplates")

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
	options.SetSort(bson.M{"author": 1})
	options.SetSkip(int64(pageToRetrieve.PageOffset))
	options.SetLimit(int64(pageToRetrieve.PageSize))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	result := make([]*models.NutritionLifestyleTemplate, 0, pageToRetrieve.PageSize)

	for cursor.Next(ctx) {
		var mongoNutritionLifestyleTemplate models.NutritionLifestyleTemplateMongoDB
		err := cursor.Decode(&mongoNutritionLifestyleTemplate)
		if err != nil {
			return nil, fmt.Errorf("cursor error: %v", err)
		}
		nutritionLifestyleTemplate, err := mongoNutritionLifestyleTemplate.ConvertToProtoMessage()
		if err != nil {
			return nil, fmt.Errorf("error parsing mongo NutritionLifestyleTemplate: %v", err)
		}
		result = append(result, nutritionLifestyleTemplate)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	if len(result) < int(req.PageSize) {
		return &pb.ListNutritionLifestyleTemplatesResponse{
			NutritionLifestyleTemplates: result,
		}, nil
	}

	currentPageToken, err := pageToRetrieve.GetToken()
	if err != nil {
		return nil, fmt.Errorf("error when forming page token: %v", err)
	}

	return &pb.ListNutritionLifestyleTemplatesResponse{
		NutritionLifestyleTemplates: result,
		RetrievedPageToken:          wrapperspb.String(currentPageToken),
	}, nil
}
