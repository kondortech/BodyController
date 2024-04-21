package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kirvader/BodyController/services/nutrition/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewNutritionClient(conn)

	resp, err := client.CreateIngredient(context.Background(), &pb.CreateIngredientRequest{
		Entity: &pb.Ingredient{
			Name: "Cucumber",
			MacrosForWeight: &pb.MacrosForWeight{
				Macros: &pb.Macros{
					Calories: 10,
					Proteins: 0,
					Carbs:    2,
					Fats:     0,
				},
				Gramms: 100,
			},
		},
	})
	// resp, err := client.DeleteIntake(context.Background(), &pb.DeleteIntakeRequest{
	// 	Id: "65edd81ff5694c28f4b0d149",
	// })
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
