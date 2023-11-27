package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	userModels "github.com/kirvader/BodyController/domains/users/models"
	pb "github.com/kirvader/BodyController/domains/users/services/aggregation/proto"
	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
	"github.com/kirvader/BodyController/pkg/utils"
)

type UsersService struct {
	authClient pbAuth.AuthClient

	pb.UnimplementedUsersServer
}

func main() {
	authServiceClientPort := utils.GetEnvWithDefault("BASE_AUTH_IP", "0.0.0.0")
	authServiceClientIP := utils.GetEnvWithDefault("BASE_AUTH_PORT", "8080")

	authServiceURI := fmt.Sprintf("%s:%s", authServiceClientPort, authServiceClientIP)
	log.Printf("uri: %s", authServiceURI)

	conn, err := grpc.Dial(authServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	authClient := pbAuth.NewAuthClient(conn)

	_, err = authClient.CreateUser(context.Background(), &pbAuth.CreateUserRequest{
		UserCredentials: &userModels.UserCredentials{
			Username: "kk-aggr",
			Password: "lol",
		},
	})
	if err != nil {
		log.Print("error when creating user: ", err)
	}

	svc := &UsersService{
		authClient: authClient,
	}

	servicePort := utils.GetEnvWithDefault("SERVICE_PORT", "10000")
	serviceURI := fmt.Sprintf(":%s", servicePort)
	log.Println("service uri: ", serviceURI)

	listener, err := net.Listen("tcp", serviceURI)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUsersServer(grpcServer, svc)
	reflection.Register(grpcServer)
	log.Printf("server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
