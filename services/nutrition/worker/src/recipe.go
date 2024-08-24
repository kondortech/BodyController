package src

import (
	"context"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/encoding/protojson"

	mongoNutrition "github.com/kirvader/BodyController/services/nutrition/mongo"
	pbNutrition "github.com/kirvader/BodyController/services/nutrition/proto"
)

func ProcessRecipeOperation(ctx context.Context, mongoClient *mongo.Client, item amqp.Delivery) error {
	switch item.Type {
	case "CREATE":
		return CreateRecipe(ctx, mongoClient, item.Body)
	case "DELETE":
		return DeleteRecipe(ctx, mongoClient, item.Body)
	}
	return nil
}

func CreateRecipe(ctx context.Context, mongoClient *mongo.Client, protoBytes []byte) error {
	var createRequest pbNutrition.CreateRecipeRequest
	err := protojson.Unmarshal(protoBytes, &createRequest)
	if err != nil {
		return fmt.Errorf("failed to parse the entity: %v", err)
	}
	log.Printf("got create request: %s", protojson.Format(&createRequest))

	coll := mongoClient.Database("BodyController").Collection("Recipes")

	mongoInstance, err := mongoNutrition.RecipeFromProto(createRequest.GetEntity())
	if err != nil {
		return fmt.Errorf("failed to convert to mongo: %v", err)
	}

	_, err = coll.InsertOne(ctx, mongoInstance)
	if err != nil {
		return fmt.Errorf("failed to insert record: %v", err)
	}

	return err
}

func DeleteRecipe(ctx context.Context, mongoClient *mongo.Client, protoBytes []byte) error {
	coll := mongoClient.Database("BodyController").Collection("Recipes")

	var deleteRequest pbNutrition.DeleteRecipeRequest
	err := protojson.Unmarshal(protoBytes, &deleteRequest)
	if err != nil {
		return err
	}
	log.Printf("got delete request: %s", protojson.Format(&deleteRequest))

	objectId, err := primitive.ObjectIDFromHex(deleteRequest.GetEntityId())
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	_, err = coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return fmt.Errorf("delete from mongo failed: %v", err)
	}
	return nil
}
