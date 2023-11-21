package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
)

// TODO replace with env pulling
const mongoDBURI = "mongodb://body-controller-mongo-db:27017"

// const mongoDBURI = "mongodb://127.0.0.1:27017"

type AuthService struct {
	mongoClient *mongo.Client

	pbAuth.UnimplementedAuthServer
}

// DeleteUser implements proto.AuthServer.
func (*AuthService) DeleteUser(context.Context, *pbAuth.DeleteUserRequest) (*pbAuth.DeleteUserResponse, error) {
	panic("unimplemented")
}

// LogIn implements proto.AuthServer.
func (*AuthService) LogIn(context.Context, *pbAuth.LogInRequest) (*pbAuth.LogInResponse, error) {
	panic("unimplemented")
}

// LogOut implements proto.AuthServer.
func (*AuthService) LogOut(context.Context, *pbAuth.LogOutRequest) (*pbAuth.LogOutResponse, error) {
	panic("unimplemented")
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBURI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	svc := &AuthService{
		mongoClient: mongoClient,
	}

	fmt.Println("Connected to MongoDB!")

	pbAuth.RegisterAuthServer(grpcServer, svc)
	reflection.Register(grpcServer)
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
