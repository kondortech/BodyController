package src

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/domains/nutrition/services/aggregation/proto"
	pbIngredient "github.com/kirvader/BodyController/domains/nutrition/services/base/ingredient/proto"

	"github.com/kirvader/BodyController/pkg/utils"
)

type NutritionService struct {
	ingredientServiceClient pbIngredient.IngredientClient

	pb.UnimplementedNutritionServer
}

func NewNutritionService(ctx context.Context) (*NutritionService, func(), error) {
	ingredientServiceClientPort := utils.GetEnvWithDefault("BASE_INGREDIENT_IP", "0.0.0.0")
	ingredientServiceClientIP := utils.GetEnvWithDefault("BASE_INGREDIENT_PORT", "20001")

	ingredientServiceURI := fmt.Sprintf("%s:%s", ingredientServiceClientPort, ingredientServiceClientIP)
	log.Printf("base-ingredient uri: %s", ingredientServiceURI)

	conn, err := grpc.Dial(ingredientServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &NutritionService{
			ingredientServiceClient: pbIngredient.NewIngredientClient(conn),
		}, func() {
			conn.Close()
		}, nil
}

func (svc *NutritionService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterNutritionServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
