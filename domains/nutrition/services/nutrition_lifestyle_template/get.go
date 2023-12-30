package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetNutritionLifestyleTemplateRequest struct {
	HexId string
}

type GetNutritionLifestyleTemplateResponse struct {
	NutritionLifestyleTemplate *models.NutritionLifestyleTemplate
}

func (svc *NutritionLifestyleTemplateService) GetIngredient(ctx context.Context, req *GetNutritionLifestyleTemplateRequest) (*GetNutritionLifestyleTemplateResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("NutritionLifestyleTemplates")

	objectId, err := primitive.ObjectIDFromHex(req.HexId)
	if err != nil {
		return nil, err
	}

	var mongo models.NutritionLifestyleTemplateMongoDB
	err = coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		options.FindOne()).
		Decode(&mongo)
	if err != nil {
		return nil, err
	}

	proto, err := mongo.ConvertToProtoMessage()
	if err != nil {
		return nil, err
	}

	return &GetNutritionLifestyleTemplateResponse{
		NutritionLifestyleTemplate: proto,
	}, nil
}
