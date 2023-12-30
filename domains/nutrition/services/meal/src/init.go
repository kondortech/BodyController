package src

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/meal/proto"
	"github.com/kirvader/BodyController/internal/db"
)

type MealService struct {
	mongoClient *mongo.Client

	pb.UnimplementedMealServer
}

func NewMealService(ctx context.Context) (*MealService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		panic(err)
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	return &MealService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}

func (svc *MealService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterMealServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
