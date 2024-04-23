package src

import (
	"context"
	"fmt"
	"net"

	"github.com/kirvader/BodyController/internal/db"
	"github.com/kirvader/BodyController/pkg/utils"
	pb "github.com/kirvader/BodyController/services/nutrition/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	mongoClient *mongo.Client

	pb.UnimplementedNutritionServer
}

func Serve(ctx context.Context) error {
	mongoClient, closeMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		return err
	}
	defer closeMongoClient()

	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		return err
	}

	servicePort := utils.GetEnvWithDefault("SERVICE_PORT", "50051")
	serviceURI := fmt.Sprintf(":%s", servicePort)
	listener, err := net.Listen("tcp", serviceURI)
	if err != nil {
		panic(err)
	}

	svc := &Service{
		mongoClient: mongoClient,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterNutritionServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
