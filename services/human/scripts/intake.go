package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/kirvader/BodyController/services/human/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewHumanClient(conn)

	resp, err := client.CreateIntake(context.Background(), &pb.CreateIntakeRequest{
		Instance: &pb.Intake{
			Macros: &pb.Macros{
				Calories: 500,
				Proteins: 40,
				Carbs:    80,
				Fats:     10,
			},
			Timestamp: &timestamppb.Timestamp{
				Seconds: time.Now().Unix(),
			},
			Username: "kondor",
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
