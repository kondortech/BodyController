package src

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/kirvader/BodyController/services/nutrition/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type mealMongo struct {
	Id                  primitive.ObjectID        `bson:"_id,omitempty"`
	Username            string                    `bson:"username"`
	RecipeId            primitive.ObjectID        `bson:"recipe_id,omitempty"`
	WeightedIngredients []weightedIngredientMongo `bson:"weighted_ingredients"`
	Timestamp           int64                     `bson:"timestamp"`
	MealStatus          uint8                     `bson:"meal_status"`
}

// CreateMeal implements proto.NutritionServer.
func (svc *Service) CreateMeal(ctx context.Context, req *pb.CreateMealRequest) (*pb.CreateMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	ingredients := make([]weightedIngredientMongo, 0, len(req.GetEntity().GetWeightedIngredients()))
	for _, item := range req.GetEntity().GetWeightedIngredients() {
		ingredients = append(ingredients, convertweightedIngredientProtoToMongo(item))
	}
	var recipeId primitive.ObjectID

	if req.GetEntity().GetRecipeId().GetValue() == "" {
		recipeId = primitive.NilObjectID
	} else {
		parsedRecipeId, err := primitive.ObjectIDFromHex(req.GetEntity().GetRecipeId().GetValue())
		if err != nil {
			return nil, err
		}
		recipeId = parsedRecipeId
	}

	resp, err := coll.InsertOne(ctx, mealMongo{
		Username:            req.GetEntity().GetUsername(),
		RecipeId:            recipeId,
		WeightedIngredients: ingredients,
		Timestamp:           req.GetEntity().GetTimestamp().GetSeconds(),
		MealStatus:          uint8(req.GetEntity().GetMealStatus()),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateMealResponse{
		EntityId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

// DeleteMeal implements proto.NutritionServer.
func (svc *Service) DeleteMeal(ctx context.Context, req *pb.DeleteMealRequest) (*pb.DeleteMealResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

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
	return &pb.DeleteMealResponse{}, nil
}

// ListMeals implements proto.NutritionServer.
func (svc *Service) ListMeals(ctx context.Context, req *pb.ListMealRequest) (*pb.ListMealsResponse, error) {
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

	coll := svc.mongoClient.Database("BodyController").Collection("Meals")

	options := options.Find()
	options.SetSort(bson.M{"timestamp": -1})
	options.SetSkip(int64(pageOffset))
	options.SetLimit(int64(pageSize))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]*pb.Meal, 0, req.GetPageSize())

	for cursor.Next(ctx) {
		var mongoInstance mealMongo
		err := cursor.Decode(&mongoInstance)
		if err != nil {
			return nil, fmt.Errorf("cursor decode error: %v", err)
		}

		ingredients := make([]*pb.WeightedIngredient, 0, len(mongoInstance.WeightedIngredients))
		for _, item := range mongoInstance.WeightedIngredients {
			ingredients = append(ingredients, convertweightedIngredientMongoToProto(item))
		}

		result = append(result, &pb.Meal{
			Id:       mongoInstance.Id.Hex(),
			Username: mongoInstance.Username,
			RecipeId: &wrapperspb.StringValue{
				Value: mongoInstance.RecipeId.Hex(),
			},
			WeightedIngredients: ingredients,
			Timestamp: &timestamppb.Timestamp{
				Seconds: mongoInstance.Timestamp,
			},
			MealStatus: pb.MealStatus(mongoInstance.MealStatus),
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	nextPageToken, err := constructPageToken(pageSize, pageOffset)
	if err != nil {
		return nil, err
	}

	return &pb.ListMealsResponse{
		Entities:      result,
		NextPageToken: nextPageToken,
	}, nil
}
