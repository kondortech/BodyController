package source

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

func ProcessMealOperation(ctx context.Context, mongoClient *mongo.Client, item amqp.Delivery) error {
	switch item.Type {
	case "CREATE":
		return CreateMeal(ctx, mongoClient, item.Body)
	case "DELETE":
		return DeleteMeal(ctx, mongoClient, item.Body)
	}
	return nil
}

func CreateMeal(ctx context.Context, mongoClient *mongo.Client, protoBytes []byte) error {
	var createRequest pbNutrition.CreateMealRequest
	err := protojson.Unmarshal(protoBytes, &createRequest)
	if err != nil {
		return err
	}
	log.Printf("got create request: %s", protojson.Format(&createRequest))

	coll := mongoClient.Database("BodyController").Collection("Meals")

	mongoInstance, err := mongoNutrition.MealFromProto(createRequest.GetEntity())
	if err != nil {
		return fmt.Errorf("parsing error: %v", err)
	}

	_, err = coll.InsertOne(ctx, mongoInstance)
	return err
}

func DeleteMeal(ctx context.Context, mongoClient *mongo.Client, protoBytes []byte) error {
	coll := mongoClient.Database("BodyController").Collection("Meals")

	var deleteRequest pbNutrition.DeleteMealRequest
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
		return fmt.Errorf("mongo execution failed: %v", err)
	}
	return err
}
