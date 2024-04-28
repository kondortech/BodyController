package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pbModels "github.com/kirvader/BodyController/services/nutrition/models"
	pb "github.com/kirvader/BodyController/services/nutrition/proto"
)

func main() {
	conn, err := grpc.Dial("0.0.0.0:50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewNutritionClient(conn)

	// resp, err := client.CreateIngredient(context.Background(), &pb.CreateIngredientRequest{
	// 	Entity: &pbModels.Ingredient{
	// 		Name: "Cucumber",
	// 		MacrosForWeight: &pbModels.MacrosForWeight{
	// 			Macros: &pbModels.Macros{
	// 				Calories: 10,
	// 				Proteins: 0,
	// 				Carbs:    2,
	// 				Fats:     0,
	// 			},
	// 			Gramms: 100,
	// 		},
	// 	},
	// })

	// resp, err := client.CreateRecipe(context.Background(), &pb.CreateRecipeRequest{
	// 	Entity: &pbModels.Recipe{
	// 		Name:        "Normal recipe 2",
	// 		RecipeSteps: "just do it",
	// 		ExampleIngredientsProportions: []*pbModels.WeightedIngredient{
	// 			{
	// 				Ingredient: &pbModels.Ingredient{
	// 					Id:   primitive.NewObjectID().Hex(),
	// 					Name: "Some good shit 2",
	// 					MacrosForWeight: &pbModels.MacrosForWeight{
	// 						Macros: &pbModels.Macros{
	// 							Calories: 10,
	// 							Proteins: 0,
	// 							Carbs:    2,
	// 							Fats:     0,
	// 						},
	// 						Gramms: 100,
	// 					},
	// 				},
	// 				Gramms: 1000,
	// 			},
	// 		},
	// 	},
	// })
	// _, err = client.DeleteRecipe(context.Background(), &pb.DeleteRecipeRequest{
	// 	EntityId: "662d4f93bd17f749e0aea14e",
	// })

	resp, err := client.CreateMeal(context.Background(), &pb.CreateMealRequest{
		Entity: &pbModels.Meal{
			Username: "Kirill",
			RecipeId: &wrapperspb.StringValue{
				Value: "662d4fffbd17f749e0aea14f",
			},
			WeightedIngredients: []*pbModels.WeightedIngredient{
				{
					Ingredient: &pbModels.Ingredient{
						Id:   primitive.NewObjectID().Hex(),
						Name: "Some good shit 2",
						MacrosForWeight: &pbModels.MacrosForWeight{
							Macros: &pbModels.Macros{
								Calories: 10,
								Proteins: 0,
								Carbs:    2,
								Fats:     0,
							},
							Gramms: 100,
						},
					},
					Gramms: 1000,
				},
			},
			Timestamp:  timestamppb.Now(),
			MealStatus: pbModels.MealStatus_Consumed,
		},
	})
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	log.Printf("response got: %v", resp)
}
