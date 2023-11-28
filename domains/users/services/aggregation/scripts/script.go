package main

import (
	"context"
	"log"

	userModels "github.com/kirvader/BodyController/domains/users/models"
	pbUsers "github.com/kirvader/BodyController/domains/users/services/aggregation/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:11000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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
