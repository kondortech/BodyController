package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kirvader/BodyController/services/nutrition/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			fmt.Println("could not close connection: ", err)
		}
	}()
	client := pb.NewNutritionClient(conn)

	resp, err := client.ListIngredients(context.Background(), &pb.ListIngredientsRequest{
		PageSize:  20,
		PageToken: nil,
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)

	id := primitive.NewObjectID()
	resp2, err2 := client.CreateIngredient(context.Background(), &pb.CreateIngredientRequest{
		Entity: &pb.Ingredient{
			Id:        id.Hex(),
			Name:      "cucumber",
			ImagePath: nil,
			MacrosNormalizedTo100G: &pb.Macros{
				Calories: 10,
				Proteins: 2,
				Carbs:    2,
				Fats:     2.3,
			},
		},
	})

	if err2 != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp2)
}
