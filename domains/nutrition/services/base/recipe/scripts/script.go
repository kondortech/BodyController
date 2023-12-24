package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:20002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pbIngredient.NewIngredientClient(conn)

	resp, err := client.ListIngredients(context.Background(), &pbIngredient.ListIngredientsRequest{
		PageSize: 10,
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
