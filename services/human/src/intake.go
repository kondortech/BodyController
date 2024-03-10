package src

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/kirvader/BodyController/services/human/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type macrosMongo struct {
	Proteins float32 `bson:"proteins"`
	Carbs    float32 `bson:"carbs"`
	Fats     float32 `bson:"fats"`
	Calories float32 `bson:"calories"`
}

type intakeMongo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Macros    macrosMongo        `bson:"macros"`
	Timestamp int64              `bson:"timestamp"`
	Username  string             `bson:"username"`
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

func (svc *Service) CreateIntake(ctx context.Context, req *pb.CreateIntakeRequest) (*pb.CreateIntakeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Intakes")

	// TODO add username validation
	resp, err := coll.InsertOne(ctx, intakeMongo{
		Macros:    convertMacrosProtoToMongo(req.GetInstance().GetMacros()),
		Timestamp: req.GetInstance().GetTimestamp().AsTime().Unix(),
		Username:  req.GetInstance().GetUsername(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateIntakeResponse{
		Id: resp.InsertedID.(primitive.ObjectID).Hex(),
	}, nil
}

func (svc *Service) DeleteIntake(ctx context.Context, req *pb.DeleteIntakeRequest) (*pb.DeleteIntakeResponse, error) {
	coll := svc.mongoClient.Database("BodyController").Collection("Intakes")

	objectId, err := primitive.ObjectIDFromHex(req.GetId())
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
	return &pb.DeleteIntakeResponse{}, nil
}

func (svc *Service) ListIntakes(ctx context.Context, req *pb.ListIntakesRequest) (*pb.ListIntakesResponse, error) {
	if req.PageSize <= 0 {
		req.PageSize = 20
	}

	coll := svc.mongoClient.Database("BodyController").Collection("Intakes")

	options := options.Find()
	// TODO add filters, maybe also by taste - so waiting
	options.SetSort(bson.M{"timestamp": 1})
	options.SetSkip(int64((req.GetPageNumber() - 1) * req.GetPageSize()))
	options.SetLimit(int64(req.GetPageSize()))

	cursor, err := coll.Find(ctx, bson.M{}, options)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	result := make([]*pb.Intake, 0, req.GetPageSize())

	for cursor.Next(ctx) {
		var mongoInstance intakeMongo
		err := cursor.Decode(&mongoInstance)
		if err != nil {
			return nil, fmt.Errorf("cursor decode error: %v", err)
		}

		result = append(result, &pb.Intake{
			Id:     mongoInstance.Id.Hex(),
			Macros: convertMacrosMongoToProto(mongoInstance.Macros),
			Timestamp: &timestamppb.Timestamp{
				Seconds: mongoInstance.Timestamp,
			},
			Username: mongoInstance.Username,
		})
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	return &pb.ListIntakesResponse{
		Instances: result,
	}, nil
}

func (svc *Service) UpdateIntake(ctx context.Context, req *pb.UpdateIntakeRequest) (*pb.UpdateIntakeResponse, error) {
	return nil, nil
}
