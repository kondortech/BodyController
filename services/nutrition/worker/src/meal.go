package src

import (
	"context"
	"fmt"
	"log"

	pb "github.com/kirvader/BodyController/services/nutrition/proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/encoding/protojson"
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
	var createRequest pb.CreateMealRequest
	err := protojson.Unmarshal(protoBytes, &createRequest)
	if err != nil {
		return err
	}
	log.Printf("got create request: %s", protojson.Format(&createRequest))

	coll := mongoClient.Database("BodyController").Collection("Meals")

	mongoInstance, err := createRequest.GetEntity().Mongo()
	if err != nil {
		return fmt.Errorf("parsing error: %v", err)
	}

	_, err = coll.InsertOne(ctx, mongoInstance)
	return err
}

func DeleteMeal(ctx context.Context, mongoClient *mongo.Client, protoBytes []byte) error {
	coll := mongoClient.Database("BodyController").Collection("Meals")

	var deleteRequest pb.DeleteMealRequest
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
