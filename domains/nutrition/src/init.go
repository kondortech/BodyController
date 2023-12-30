package src

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/kirvader/BodyController/domains/nutrition/proto"
	ingredient "github.com/kirvader/BodyController/domains/nutrition/services/ingredient"
	meal "github.com/kirvader/BodyController/domains/nutrition/services/meal"
	nutritionLifestyleTemplate "github.com/kirvader/BodyController/domains/nutrition/services/nutrition_lifestyle_template"
	personalNutritionLifestyle "github.com/kirvader/BodyController/domains/nutrition/services/personal_nutrition_lifestyle"
	recipe "github.com/kirvader/BodyController/domains/nutrition/services/recipe"
)

type NutritionService struct {
	ingredientService                 *ingredient.Service
	recipeService                     *recipe.Service
	nutritionLifestyleTemplateService *nutritionLifestyleTemplate.Service
	personalNutritionLifestyleService *personalNutritionLifestyle.Service
	mealService                       *meal.Service

	pb.UnimplementedNutritionServer
}

func NewNutritionService(ctx context.Context) (*NutritionService, func(), error) {
	closeFunctions := make([]func(), 0)

	ingredientService, ingredientCloseFunc, err := ingredient.NewService(ctx)
	closeFunctions = append(closeFunctions, ingredientCloseFunc)
	if err != nil {
		for _, closeFunc := range closeFunctions {
			closeFunc()
		}

		return nil, func() {}, fmt.Errorf("ingredient service initialization error: %v", err)
	}

	recipeService, recipeCloseFunc, err := recipe.NewService(ctx)
	closeFunctions = append(closeFunctions, recipeCloseFunc)
	if err != nil {
		for _, closeFunc := range closeFunctions {
			closeFunc()
		}

		return nil, func() {}, fmt.Errorf("recipe service initialization error: %v", err)
	}

	nutritionLifestyleTemplateService, nutritionLifestyleTemplateCloseFunc, err := nutritionLifestyleTemplate.NewService(ctx)
	closeFunctions = append(closeFunctions, nutritionLifestyleTemplateCloseFunc)
	if err != nil {
		for _, closeFunc := range closeFunctions {
			closeFunc()
		}

		return nil, func() {}, fmt.Errorf("nutrition lifestyle template service initialization error: %v", err)
	}

	personalNutritionLifestyleService, personalNutritionLifestyleCloseFunc, err := personalNutritionLifestyle.NewService(ctx)
	closeFunctions = append(closeFunctions, personalNutritionLifestyleCloseFunc)
	if err != nil {
		for _, closeFunc := range closeFunctions {
			closeFunc()
		}

		return nil, func() {}, fmt.Errorf("personal nutrition lifestyle service initialization error: %v", err)
	}

	mealService, mealCloseFunc, err := meal.NewService(ctx)
	closeFunctions = append(closeFunctions, mealCloseFunc)
	if err != nil {
		for _, closeFunc := range closeFunctions {
			closeFunc()
		}

		return nil, func() {}, fmt.Errorf("recipe client initialization error: %v", err)
	}

	return &NutritionService{
			ingredientService:                 ingredientService,
			recipeService:                     recipeService,
			nutritionLifestyleTemplateService: nutritionLifestyleTemplateService,
			personalNutritionLifestyleService: personalNutritionLifestyleService,
			mealService:                       mealService,
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
