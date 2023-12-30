package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/meal/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:20005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMealClient(conn)

	resp, err := client.ListMeals(context.Background(), &pb.ListMealsRequest{
		PageSize: 10,
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
