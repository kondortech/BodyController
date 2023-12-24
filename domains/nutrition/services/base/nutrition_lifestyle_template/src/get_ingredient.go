package src

import (
	"context"

	"github.com/kirvader/BodyController/domains/nutrition/models"
	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/nutrition_lifestyle_template/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (svc *NutritionLifestyleTemplateService) GetIngredient(ctx context.Context, req *pb.GetNutritionLifestyleTemplateRequest) (*pb.GetNutritionLifestyleTemplateResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("NutritionLifestyleTemplates")

	objectId, err := primitive.ObjectIDFromHex(req.NutritionLifestyleTemplateHexId)
	if err != nil {
		return nil, err
	}

	var nutritionLifestyleTemplate models.NutritionLifestyleTemplateMongoDB
	err = coll.FindOne(context.TODO(),
		bson.D{{Key: "_id", Value: objectId}},
		options.FindOne()).
		Decode(&nutritionLifestyleTemplate)
	if err != nil {
		return nil, err
	}

	result, err := nutritionLifestyleTemplate.ConvertToProtoMessage()
	if err != nil {
		return nil, err
	}

	return &pb.GetNutritionLifestyleTemplateResponse{
		NutritionLifestyleTemplate: result,
	}, nil
}
