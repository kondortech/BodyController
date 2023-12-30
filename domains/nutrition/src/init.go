package src

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/domains/nutrition/proto"
	ingredient "github.com/kirvader/BodyController/domains/nutrition/services/ingredient"
	recipe "github.com/kirvader/BodyController/domains/nutrition/services/recipe"

	"github.com/kirvader/BodyController/pkg/utils"
)

type NutritionService struct {
	ingredientService *ingredient.IngredientService
	recipeService     *recipe.RecipeService

	pb.UnimplementedNutritionServer
}

func NewNutritionService(ctx context.Context) (*NutritionService, func(), error) {
	ingredientServiceClientPort := utils.GetEnvWithDefault("BASE_INGREDIENT_IP", "0.0.0.0")
	ingredientServiceClientIP := utils.GetEnvWithDefault("BASE_INGREDIENT_PORT", "20001")

	ingredientServiceURI := fmt.Sprintf("%s:%s", ingredientServiceClientPort, ingredientServiceClientIP)
	log.Printf("base-ingredient uri: %s", ingredientServiceURI)

	closeFunctions := make([]func(), 0)

	conn, err := grpc.Dial(ingredientServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	closeFunctions = append(closeFunctions, func() {
		conn.Close()
	})

	ingredientService, ingredientDBConnectionClose, err := ingredient.NewIngredientService(ctx)
	closeFunctions = append(closeFunctions, ingredientDBConnectionClose)
	if err != nil {
		for _, closeFunc := range closeFunctions {
			closeFunc()
		}

		return nil, func() {}, fmt.Errorf("ingredient client initialization error: %v", err)
	}

	recipeService, recipeDBConnectionClose, err := recipe.NewRecipeService(ctx)
	closeFunctions = append(closeFunctions, recipeDBConnectionClose)
	if err != nil {
		for _, closeFunc := range closeFunctions {
			closeFunc()
		}

		return nil, func() {}, fmt.Errorf("recipe client initialization error: %v", err)
	}

	return &NutritionService{
			ingredientService: ingredientService,
			recipeService:     recipeService,
		}, func() {
			for _, closeFunc := range closeFunctions {
				closeFunc()
			}
		}, nil
}

func (svc *NutritionService) Serve(listener net.Listener) error {
	grpcServer := grpc.NewServer()
	pb.RegisterNutritionServer(grpcServer, svc)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener)
}
