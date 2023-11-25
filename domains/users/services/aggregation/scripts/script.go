package main

import (
	"context"
	"log"

	pbUsers "github.com/kirvader/BodyController/domains/users/services/aggregation/proto"
	userModels "github.com/kirvader/BodyController/models/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbUsers.NewUsersClient(conn)

	resp, err := client.CreateUser(context.Background(), &pbUsers.CreateUserRequest{
		UserCredentials: &userModels.UserCredentials{
			Username: "kk-aggr",
			Password: "lol",
		},
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
