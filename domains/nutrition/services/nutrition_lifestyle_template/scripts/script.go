package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/base/nutrition_lifestyle_template/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:2000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewNutritionLifestyleTemplateClient(conn)

	resp, err := client.ListNutritionLifestyleTemplates(context.Background(), &pb.ListNutritionLifestyleTemplatesRequest{
		PageSize: 10,
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
