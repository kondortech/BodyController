package main

import (
	"context"
	"log"

	pbAuth "github.com/kirvader/BodyController/domains/users/services/base/auth/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:10001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbAuth.NewAuthClient(conn)

	resp, err := client.DeleteUser(context.Background(), &pbAuth.DeleteUserRequest{
		Username: "kk2",
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
