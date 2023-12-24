package models

import (
	"errors"
	"fmt"

	pbTypes "github.com/kirvader/BodyController/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EncodeMealStatus(mealStatus MealStatus) string {
	return mealStatus.String()
}

func DecodeMealStatus(encodedMealString string) (MealStatus, error) {
	if intValue, ok := MealStatus_value[encodedMealString]; ok {
		return MealStatus(intValue), nil
	}
	return MealStatus_NOT_SET, fmt.Errorf("DecodeMealStatus: MealStatus couldn't decode %q", encodedMealString)

}

type NutritionStepRefMongo struct {
	ExecutionCycleNumber int64 `bson:"ExecutionCycleNumber"`
	NutritionStepIndex   int64 `bson:"NutritionStepIndex"`
}

func (proto *NutritionStepRef) ConvertToMongoDocument() (*NutritionStepRefMongo, error) {
	return &NutritionStepRefMongo{
		ExecutionCycleNumber: proto.ExecutionCycleNumber,
		NutritionStepIndex:   proto.NutritionStepIndex,
	}, nil
}

func (mongo *NutritionStepRefMongo) ConvertToProtoMessage() (*NutritionStepRef, error) {
	return &NutritionStepRef{
		ExecutionCycleNumber: mongo.ExecutionCycleNumber,
		NutritionStepIndex:   mongo.NutritionStepIndex,
	}, nil
}

type MealMongo struct {
	Username                     string
	NutritionStepRef             *NutritionStepRefMongo
	PersonalNutritionLifestyleId primitive.ObjectID
	StatusUpdateTime             *pbTypes.DateTimeMongo
	MealStatus                   string

	RecipeId                primitive.ObjectID
	MealRequiredIngredients []*WeightedIngredientMongoDB
	Macros                  *MacrosMongoDB
}

func (proto *Meal) ConvertToMongoDocument() (*MealMongo, error) {
	nutritionStepRefMongo, _ := proto.NutritionStepRef.ConvertToMongoDocument()
	macrosMongo, _ := proto.Macros.ConvertToMongoDocument()
	statusUpdateTime, _ := proto.StatusUpdateTime.ConvertToMongoDocument()

	if len(proto.GetRecipeId()) == 0 {
		return nil, errors.New("Meal.ConvertToMongoDocument: proto.RecipeId is empty")
	}
	recipeId, err := primitive.ObjectIDFromHex(proto.GetRecipeId())
	if err != nil {
		return nil, fmt.Errorf("Meal.ConvertToMongoDocument: returned error when parsing recipe_id: %v", err)
	}

	if len(proto.GetPersonalNutritionLifestyleId()) == 0 {
		return nil, errors.New("Meal.ConvertToMongoDocument: proto.PersonalNutritionLifestyleId is empty")
	}
	personalNutritionLifestyleId, err := primitive.ObjectIDFromHex(proto.GetPersonalNutritionLifestyleId())
	if err != nil {
		return nil, fmt.Errorf("Meal.ConvertToMongoDocument: returned error when parsing personal_nutrition_lifestyle_id: %v", err)
	}
	requiredIngredients := make([]*WeightedIngredientMongoDB, 0, len(proto.MealRequiredIngredients))

	for _, ingr := range proto.MealRequiredIngredients {
		mongoIngr, _ := ingr.ConvertToMongoDocument()
		requiredIngredients = append(requiredIngredients, mongoIngr)
	}

	return &MealMongo{
		Username:                     proto.Username,
		NutritionStepRef:             nutritionStepRefMongo,
		PersonalNutritionLifestyleId: personalNutritionLifestyleId,
		StatusUpdateTime:             statusUpdateTime,
		MealStatus:                   EncodeMealStatus(proto.MealStatus),
		RecipeId:                     recipeId,
		MealRequiredIngredients:      requiredIngredients,
		Macros:                       macrosMongo,
	}, nil
}

func (mongo *MealMongo) ConvertToProtoMessage() (*Meal, error) {
	mealStatus, err := DecodeMealStatus(mongo.MealStatus)
	if err != nil {
		return nil, err
	}
	nutritionStepRefProto, _ := mongo.NutritionStepRef.ConvertToProtoMessage()
	macrosProto, _ := mongo.Macros.ConvertToProtoMessage()
	statusUpdateTimeProto, _ := mongo.StatusUpdateTime.ConvertToProtoMessage()

	requiredIngredientsProto := make([]*WeightedIngredient, 0, len(mongo.MealRequiredIngredients))
	for _, ingr := range mongo.MealRequiredIngredients {
		protoIngr, _ := ingr.ConvertToProtoMessage()
		requiredIngredientsProto = append(requiredIngredientsProto, protoIngr)
	}

	return &Meal{
		Username:                     mongo.Username,
		NutritionStepRef:             nutritionStepRefProto,
		PersonalNutritionLifestyleId: mongo.PersonalNutritionLifestyleId.Hex(),
		StatusUpdateTime:             statusUpdateTimeProto,
		MealStatus:                   mealStatus,
		RecipeId:                     mongo.RecipeId.Hex(),
		MealRequiredIngredients:      requiredIngredientsProto,
		Macros:                       macrosProto,
	}, nil
}
