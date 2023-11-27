package src

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
	"github.com/kirvader/BodyController/internal/db"
)

type IngredientService struct {
	mongoClient *mongo.Client

	pbIngredient.UnimplementedIngredientServer
}

func NewIngredientService(ctx context.Context) (*IngredientService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		panic(err)
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	return &IngredientService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}

func (svc *IngredientService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pbIngredient.RegisterIngredientServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
