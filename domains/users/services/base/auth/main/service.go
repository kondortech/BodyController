package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
	"github.com/kirvader/BodyController/internal/db"
	"github.com/kirvader/BodyController/pkg/utils"
)

type AuthService struct {
	mongoClient *mongo.Client

	pbAuth.UnimplementedAuthServer
}

func main() {
	servicePort := utils.GetEnvWithDefault("SERVICE_PORT", "10000")
	serviceURI := fmt.Sprintf(":%s", servicePort)
	log.Println("service uri: ", serviceURI)

	listener, err := net.Listen("tcp", serviceURI)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mongoClient, disconnectMongoClient, err := db.InitMongoDBClientFromENV(ctx)
	if err != nil {
		panic(err)
	}
	defer disconnectMongoClient()
	if err = db.PingMongoDb(ctx, mongoClient); err != nil {
		panic(err)
	}

	svc := &AuthService{
		mongoClient: mongoClient,
	}

	pbAuth.RegisterAuthServer(grpcServer, svc)
	reflection.Register(grpcServer)
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
