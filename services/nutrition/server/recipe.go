package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/encoding/protojson"

	models "github.com/kirvader/BodyController/services/nutrition/models"
	pb "github.com/kirvader/BodyController/services/nutrition/proto"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (svc *Service) CreateRecipe(ctx context.Context, req *pb.CreateRecipeRequest) (*pb.CreateRecipeResponse, error) {
	if req == nil || req.Entity == nil { // TODO add real validation
		return nil, errors.New("nil instance")
	}
	req.Entity.Id = primitive.NewObjectID().Hex()

	body, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = svc.workerChannelMQ.PublishWithContext(ctx,
		"",       // exchange
		"recipe", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Type:        OperationCreate,
			Timestamp:   time.Now(),
			Body:        []byte(body),
		})
	if err != nil {
		return nil, fmt.Errorf("failed to publish event: %s", err)
	}
	log.Println("published CREATE event with id: ", req.Entity.Id)

	return &pb.CreateRecipeResponse{
		EntityId: req.Entity.Id,
	}, nil
}

func (svc *Service) DeleteRecipe(ctx context.Context, req *pb.DeleteRecipeRequest) (*pb.DeleteRecipeResponse, error) {
	if req == nil || req.EntityId == "" { // TODO add real validation
		return nil, errors.New("nil instance")
	}

	body, err := protojson.Marshal(req)
	if err != nil {
		return nil, err
	}

	err = svc.workerChannelMQ.PublishWithContext(ctx,
		"",       // exchange
		"recipe", // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Type:        OperationDelete,
			Timestamp:   time.Now(),
			Body:        []byte(body),
		})
	if err != nil {
		return nil, fmt.Errorf("failed to publish event: %s", err)
	}
	log.Println("published DELETE event with id: ", req.EntityId)

	return &pb.DeleteRecipeResponse{}, nil
	// coll := svc.mongoClient.Database("BodyController").Collection("Recipes")

	// objectId, err := primitive.ObjectIDFromHex(req.GetEntityId())
	// if err != nil {
	// 	return nil, err
	// }

	// resp, err := coll.DeleteOne(ctx, bson.M{"_id": objectId})
	// if err != nil {
	// 	return nil, err
	// }

	// if resp.DeletedCount != 1 {
	// 	return nil, errors.New("delete operation failed")
	// }
	// return &pb.DeleteRecipeResponse{}, nil
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

	result := make([]*models.Recipe, 0, req.GetPageSize())

	for cursor.Next(ctx) {
		var mongoInstance models.RecipeMongo
		err := cursor.Decode(&mongoInstance)
		if err != nil {
			return nil, fmt.Errorf("cursor decode error: %v", err)
		}

		instance, err := mongoInstance.Proto()
		if err != nil {
			return nil, fmt.Errorf("failed to parse entity: %v", err)
		}

		result = append(result, instance)
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
