package src

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/personal_nutrition_lifestyle/proto"
	"github.com/kirvader/BodyController/internal/db"
)

type PersonalNutritionLifestyleService struct {
	mongoClient *mongo.Client

	pb.UnimplementedPersonalNutritionLifestyleServer
}

func NewPersonalNutritionLifestyleService(ctx context.Context) (*PersonalNutritionLifestyleService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		panic(err)
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	return &PersonalNutritionLifestyleService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}

func (svc *PersonalNutritionLifestyleService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterPersonalNutritionLifestyleServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
