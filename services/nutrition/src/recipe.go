package src

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/kirvader/BodyController/services/nutrition/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type weightedIngredientMongo struct {
	Ingredient ingredientMongo `bson:"ingredient"`
	Gramms     float32         `bson:"gramms"`
}

func convertweightedIngredientProtoToMongo(instance *pb.WeightedIngredient) weightedIngredientMongo {
	return weightedIngredientMongo{
		Ingredient: ingredientMongo{
			Id:              primitive.ObjectID([]byte(instance.GetIngredient().GetId())),
			Name:            instance.GetIngredient().GetName(),
			MacrosForWeight: convertMacrosForWeightProtoToMongo(instance.GetIngredient().GetMacrosForWeight()),
		},
		Gramms: instance.GetGramms(),
	}
}

func convertweightedIngredientMongoToProto(instance weightedIngredientMongo) *pb.WeightedIngredient {
	return &pb.WeightedIngredient{
		Ingredient: &pb.Ingredient{
			Id:              instance.Ingredient.Id.Hex(),
			Name:            instance.Ingredient.Name,
			MacrosForWeight: convertMacrosForWeightMongoToProto(instance.Ingredient.MacrosForWeight),
		},
		Gramms: instance.Gramms,
	}
}

type recipeMongo struct {
	Id                            primitive.ObjectID        `bson:"_id,omitempty"`
	Name                          string                    `bson:"name"`
	RecipeSteps                   string                    `bson:"recipe_steps"`
	ExampleIngredientsProportions []weightedIngredientMongo `bson:"example_ingredients_proportions"`
}

// CreateRecipe implements proto.NutritionServer.
func (svc *Service) CreateRecipe(ctx context.Context, req *pb.CreateRecipeRequest) (*pb.CreateRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	ingredients := make([]weightedIngredientMongo, 0, len(req.GetEntity().GetExampleIngredientsProportions()))
	for _, item := range req.GetEntity().GetExampleIngredientsProportions() {
		ingredients = append(ingredients, convertweightedIngredientProtoToMongo(item))
	}

	resp, err := coll.InsertOne(ctx, recipeMongo{
		Name:                          req.GetEntity().GetName(),
		RecipeSteps:                   req.GetEntity().GetRecipeSteps(),
		ExampleIngredientsProportions: ingredients,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateRecipeResponse{
		EntityId: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

// DeleteRecipe implements proto.NutritionServer.
func (svc *Service) DeleteRecipe(ctx context.Context, req *pb.DeleteRecipeRequest) (*pb.DeleteRecipeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

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
	return &pb.DeleteRecipeResponse{}, nil
}

// ListRecipes implements proto.NutritionServer.
func (svc *Service) ListRecipes(ctx context.Context, req *pb.ListRecipeRequest) (*pb.ListRecipesResponse, error) {
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

	coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	options := options.Find()
	options.SetSort(bson.M{"name": 1})
	options.SetSkip(int64(pageOffset))
	options.SetLimit(int64(pageSize))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	result := make([]*pb.Recipe, 0, req.GetPageSize())

	for cursor.Next(ctx) {
		var mongoInstance recipeMongo
		err := cursor.Decode(&mongoInstance)
		if err != nil {
			return nil, fmt.Errorf("cursor decode error: %v", err)
		}

		ingredients := make([]*pb.WeightedIngredient, 0, len(mongoInstance.ExampleIngredientsProportions))
		for _, item := range mongoInstance.ExampleIngredientsProportions {
			ingredients = append(ingredients, convertweightedIngredientMongoToProto(item))
		}

		result = append(result, &pb.Recipe{
			Id:                            mongoInstance.Id.Hex(),
			Name:                          mongoInstance.Name,
			RecipeSteps:                   mongoInstance.RecipeSteps,
			ExampleIngredientsProportions: ingredients,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	nextPageToken, err := constructPageToken(pageSize, pageOffset)
	if err != nil {
		return nil, err
	}

	return &pb.ListRecipesResponse{
		Entities:      result,
		NextPageToken: nextPageToken,
	}, nil
}
