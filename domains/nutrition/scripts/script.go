package main

import (
	"context"
	"log"

	pbNutrition "github.com/kirvader/BodyController/domains/nutrition/services/aggregation/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:20000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbNutrition.NewNutritionClient(conn)

	resp, err := client.ListIngredients(context.Background(), &pbNutrition.ListIngredientsRequest{
		PageSize: 10,
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
