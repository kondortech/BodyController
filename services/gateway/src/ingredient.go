package src

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/kirvader/BodyController/services/gateway/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type macrosMongo struct {
	Proteins float32 `bson:"proteins"`
	Carbs    float32 `bson:"carbs"`
	Fats     float32 `bson:"fats"`
	Calories float32 `bson:"calories"`
}

func convertMacrosProtoToMongo(instance *pb.Macros) macrosMongo {
	return macrosMongo{
		Proteins: float32(instance.GetProteins()),
		Carbs:    float32(instance.GetCarbs()),
		Fats:     float32(instance.GetFats()),
		Calories: float32(instance.GetCalories()),
	}
}

func convertMacrosMongoToProto(instance macrosMongo) *pb.Macros {
	return &pb.Macros{
		Calories: instance.Calories,
		Proteins: instance.Proteins,
		Carbs:    instance.Carbs,
		Fats:     instance.Fats,
	}
}

type macrosForWeightMongo struct {
	Macros macrosMongo `bson:"macros"`
	Gramms float32     `bson:"gramms"`
}

func convertMacrosForWeightProtoToMongo(instance *pb.MacrosForWeight) macrosForWeightMongo {
	return macrosForWeightMongo{
		Macros: convertMacrosProtoToMongo(instance.GetMacros()),
		Gramms: instance.GetGramms(),
	}
}

func convertMacrosForWeightMongoToProto(instance macrosForWeightMongo) *pb.MacrosForWeight {
	return &pb.MacrosForWeight{
		Macros: convertMacrosMongoToProto(instance.Macros),
		Gramms: instance.Gramms,
	}
}

type ingredientMongo struct {
	Id              primitive.ObjectID   `bson:"_id,omitempty"`
	Name            string               `bson:"name"`
	MacrosForWeight macrosForWeightMongo `bson:"macros_for_weight"`
}

func (svc *Service) CreateIngredient(ctx context.Context, req *pb.CreateIngredientRequest) (*pb.CreateIngredientResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	resp, err := coll.InsertOne(ctx, ingredientMongo{
		Name:            req.GetEntity().GetId(),
		MacrosForWeight: convertMacrosForWeightProtoToMongo(req.GetEntity().GetMacrosForWeight()),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateIngredientResponse{
		EntityId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

func (svc *Service) DeleteIngredient(ctx context.Context, req *pb.DeleteIngredientRequest) (*pb.DeleteIngredientResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	objectId, err := primitive.ObjectIDFromHex(req.GetEntityId())
	if err != nil {
		return nil, err
	}

	resp, err := coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return nil, err
	}

	if resp.DeletedCount != 1 {
		return nil, errors.New("delete operation failed")
	}
	return &pb.DeleteIngredientResponse{}, nil
}

// ListIngredients implements proto.NutritionServer.
func (svc *Service) ListIngredients(ctx context.Context, req *pb.ListIngredientRequest) (*pb.ListIngredientsResponse, error) {
	var pageSize, pageOffset int32
	if req.GetPageToken() == "" {
		pageSizeFromToken, pageOffsetFromToken, err := deconstructPageToken(req.GetPageToken())
		if err != nil {
			return nil, err
		}
		pageSize, pageOffset = pageSizeFromToken, pageOffsetFromToken
	} else {
		pageOffset = 0
		if req.GetPageSize() <= 0 {
			pageSize = 100
		} else if req.GetPageSize() >= 500 {
			pageSize = 500
		} else {
			pageSize = req.GetPageSize()
		}
	}

	coll := svc.mongoClient.Database("BodyController").Collection("Ingredients")

	options := options.Find()
	options.SetSort(bson.M{"name": 1})
	options.SetSkip(int64(pageOffset))
	options.SetLimit(int64(pageSize))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]*pb.Ingredient, 0, req.GetPageSize())

	for cursor.Next(ctx) {
		var mongoInstance ingredientMongo
		err := cursor.Decode(&mongoInstance)
		if err != nil {
			return nil, fmt.Errorf("cursor decode error: %v", err)
		}

		result = append(result, &pb.Ingredient{
			Id:              mongoInstance.Id.Hex(),
			Name:            mongoInstance.Name,
			MacrosForWeight: convertMacrosForWeightMongoToProto(mongoInstance.MacrosForWeight),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	nextPageToken, err := constructPageToken(pageSize, pageOffset)
	if err != nil {
		return nil, err
	}

	return &pb.ListIngredientsResponse{
		Entities:      result,
		NextPageToken: nextPageToken,
	}, nil
}
