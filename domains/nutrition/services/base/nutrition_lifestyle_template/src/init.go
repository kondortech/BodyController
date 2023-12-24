package src

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/nutrition_lifestyle_template/proto"
	"github.com/kirvader/BodyController/internal/db"
)

type NutritionLifestyleTemplateService struct {
	mongoClient *mongo.Client

	pb.UnimplementedNutritionLifestyleTemplateServer
}

func NewNutritionLifestyleTemplateService(ctx context.Context) (*NutritionLifestyleTemplateService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		panic(err)
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	return &NutritionLifestyleTemplateService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}

func (svc *NutritionLifestyleTemplateService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterNutritionLifestyleTemplateServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
