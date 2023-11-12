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

	pbAuth "github.com/kirvader/BodyController/services/users/auth/proto"
)

// TODO replace with env pulling
const mongoDBURI = "mongodb://body-controller-mongo-db:27017"

type AuthService struct {
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

	s := grpc.NewServer()

	reflection.Register(s)

	svc := &AuthService{}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBURI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	pbAuth.RegisterAuthServer(s, svc)
	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
