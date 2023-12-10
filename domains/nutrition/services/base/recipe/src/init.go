package src

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbRecipe "github.com/kirvader/BodyController/domains/nutrition/services/base/recipe/proto"
	"github.com/kirvader/BodyController/internal/db"
)

type RecipeService struct {
	mongoClient *mongo.Client

	pbRecipe.UnimplementedRecipeServer
}

func NewService(ctx context.Context) (*RecipeService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		panic(err)
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	return &RecipeService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}

func (svc *RecipeService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pbRecipe.RegisterRecipeServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
