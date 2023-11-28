package src

import (
	"context"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
	"github.com/kirvader/BodyController/internal/db"
)

type AuthService struct {
	mongoClient *mongo.Client

	pbAuth.UnimplementedAuthServer
}

func NewAuthService(ctx context.Context) (*AuthService, func(), error) {
	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		panic(err)
	}
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	return &AuthService{
		mongoClient: mongoClient,
	}, disconnectMongoClient, nil
}

func (svc *AuthService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pbAuth.RegisterAuthServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
