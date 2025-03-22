package mongo

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pbNutrition "github.com/kirvader/BodyController/services/nutrition/proto"
)

var ErrNilInstance = errors.New("nil instance")
var ErrInvalidId = errors.New("invalid id")

type Macros struct {
	Proteins float32 `bson:"proteins"`
	Carbs    float32 `bson:"carbs"`
	Fats     float32 `bson:"fats"`
	Calories float32 `bson:"calories"`
}

func MacrosFromProto(protoEntity *pbNutrition.Macros) (*Macros, error) {
	if protoEntity == nil {
		return nil, ErrNilInstance
	}
	return &Macros{
		Proteins: protoEntity.GetProteins(),
		Carbs:    float32(protoEntity.GetCarbs()),
		Fats:     float32(protoEntity.GetFats()),
		Calories: float32(protoEntity.GetCalories()),
	}, nil
}

func MacrosToProto(mongoEntity *Macros) (*pbNutrition.Macros, error) {
	if mongoEntity == nil {
		return nil, ErrNilInstance
	}
	return &pbNutrition.Macros{
		Calories: mongoEntity.Calories,
		Proteins: mongoEntity.Proteins,
		Carbs:    mongoEntity.Carbs,
		Fats:     mongoEntity.Fats,
	}, nil
}

type Ingredient struct {
	Id                     primitive.ObjectID `bson:"_id,omitempty"`
	Name                   string             `bson:"name"`
	ImagePath              string             `bson:"image_path"`
	MacrosNormalizedTo100g *Macros            `bson:"macros_normalized_to_100g"`
}

func IngredientFromProto(protoEntity *pbNutrition.Ingredient) (*Ingredient, error) {
	if protoEntity == nil {
		return nil, ErrNilInstance
	}

	macrosNormalizedTo100g, err := MacrosFromProto(protoEntity.GetMacrosNormalizedTo100G())
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(protoEntity.GetId())
	if err != nil {
		return nil, ErrInvalidId
	}

	var imagePath string
	if protoEntity.GetImagePath() != nil {
		imagePath = protoEntity.GetImagePath().GetValue()
	}

	return &Ingredient{
		Id:                     id,
		Name:                   protoEntity.GetName(),
		ImagePath:              imagePath,
		MacrosNormalizedTo100g: macrosNormalizedTo100g,
	}, nil
}

func IngredientToProto(mongoEntity *Ingredient) (*pbNutrition.Ingredient, error) {
	if mongoEntity == nil {
		return nil, ErrNilInstance
	}

	macrosNormalizedTo100g, err := MacrosToProto(mongoEntity.MacrosNormalizedTo100g)
	if err != nil {
		return nil, err
	}

	var imagePath *wrapperspb.StringValue
	if mongoEntity.ImagePath != "" {
		imagePath = &wrapperspb.StringValue{
			Value: mongoEntity.ImagePath,
		}
	}

	return &pbNutrition.Ingredient{
		Id:                     mongoEntity.Id.Hex(),
		Name:                   mongoEntity.Name,
		ImagePath:              imagePath,
		MacrosNormalizedTo100G: macrosNormalizedTo100g,
	}, nil
}

type WeightedIngredient struct {
	IngredientId primitive.ObjectID `bson:"ingredient_id"`
	Mass         float32            `bson:"mass"`
}

func WeightedIngredientFromProto(protoEntity *pbNutrition.WeightedIngredient) (*WeightedIngredient, error) {
	ingredientId, err := primitive.ObjectIDFromHex(protoEntity.GetIngredientId())
	if err != nil {
		return nil, ErrInvalidId
	}

	return &WeightedIngredient{
		IngredientId: ingredientId,
		Mass:         protoEntity.GetMass(),
	}, nil
}

func WeightedIngredientToProto(mongoEntity *WeightedIngredient) (*pbNutrition.WeightedIngredient, error) {
	if mongoEntity == nil {
		return nil, ErrNilInstance
	}

	return &pbNutrition.WeightedIngredient{
		IngredientId: mongoEntity.IngredientId.Hex(),
		Mass:         mongoEntity.Mass,
	}, nil
}

type Recipe struct {
	Id                            primitive.ObjectID    `bson:"_id,omitempty"`
	Name                          string                `bson:"name"`
	ImagePath                     string                `bson:"image_path"`
	TasteDescription              string                `bson:"taste_description"`
	CookingStepsDescription       string                `bson:"cooking_steps_description"`
	OriginalIngredientsProportion []*WeightedIngredient `bson:"original_ingredients_proportion"`
}

func RecipeFromProto(protoEntity *pbNutrition.Recipe) (*Recipe, error) {
	if protoEntity == nil {
		return nil, ErrNilInstance
	}

	id, err := primitive.ObjectIDFromHex(protoEntity.GetId())
	if err != nil {
		return nil, ErrInvalidId
	}

	weightedIngredients := make([]*WeightedIngredient, 0, len(protoEntity.GetOriginalIngredientsProportion()))
	for _, protoWeightedIngredient := range protoEntity.GetOriginalIngredientsProportion() {
		mongoWeightedIngedient, err := WeightedIngredientFromProto(protoWeightedIngredient)
		if err != nil {
			return nil, err
		}

		weightedIngredients = append(weightedIngredients, mongoWeightedIngedient)
	}

	var imagePath string
	if protoEntity.GetImagePath() != nil {
		imagePath = protoEntity.GetImagePath().GetValue()
	}

	return &Recipe{
		Id:                            id,
		Name:                          protoEntity.GetName(),
		ImagePath:                     imagePath,
		TasteDescription:              protoEntity.GetTasteDescription(),
		CookingStepsDescription:       protoEntity.GetCookingStepsDescription(),
		OriginalIngredientsProportion: weightedIngredients,
	}, nil
}

func RecipeToProto(mongoEntity *Recipe) (*pbNutrition.Recipe, error) {
	if mongoEntity == nil {
		return nil, ErrNilInstance
	}

	weightedIngredients := make([]*pbNutrition.WeightedIngredient, 0, len(mongoEntity.OriginalIngredientsProportion))
	for _, mongoWeightedIngredient := range mongoEntity.OriginalIngredientsProportion {
		protoWeightedIngredient, err := WeightedIngredientToProto(mongoWeightedIngredient)
		if err != nil {
			return nil, err
		}

		weightedIngredients = append(weightedIngredients, protoWeightedIngredient)
	}

	var imagePath *wrapperspb.StringValue
	if mongoEntity.ImagePath != "" {
		imagePath = &wrapperspb.StringValue{
			Value: mongoEntity.ImagePath,
		}
	}

	return &pbNutrition.Recipe{
		Id:                            mongoEntity.Id.Hex(),
		Name:                          mongoEntity.Name,
		TasteDescription:              mongoEntity.TasteDescription,
		CookingStepsDescription:       mongoEntity.CookingStepsDescription,
		OriginalIngredientsProportion: weightedIngredients,
		ImagePath:                     imagePath,
	}, nil
}

type Meal struct {
	Id                   primitive.ObjectID    `bson:"_id,omitempty"`
	Author               string                `bson:"author"`
	ImagePath            string                `bson:"image_path"`
	BaseRecipeId         primitive.ObjectID    `bson:"base_recipe_id,omitempty"`
	WeightedIngredients  []*WeightedIngredient `bson:"weighted_ingredients"`
	ConsumptionTimestamp int64                 `bson:"consumption_timestamp"`
}

func MealFromProto(protoEntity *pbNutrition.Meal) (*Meal, error) {
	if protoEntity == nil {
		return nil, ErrNilInstance
	}

	id, err := primitive.ObjectIDFromHex(protoEntity.GetId())
	if err != nil {
		return nil, ErrInvalidId
	}

	var baseRecipeId primitive.ObjectID
	if protoEntity.GetBaseRecipeId() != nil {
		baseRecipeId, err = primitive.ObjectIDFromHex(protoEntity.GetBaseRecipeId().GetValue())
		if err != nil {
			return nil, ErrInvalidId
		}
	}

	weightedIngredients := make([]*WeightedIngredient, 0, len(protoEntity.GetWeightedIngredients()))
	for _, weightedIngredient := range protoEntity.GetWeightedIngredients() {
		weightedIngredientMongo, err := WeightedIngredientFromProto(weightedIngredient)
		if err != nil {
			return nil, err
		}
		weightedIngredients = append(weightedIngredients, weightedIngredientMongo)
	}

	var imagePath string
	if protoEntity.GetImagePath() != nil {
		imagePath = protoEntity.GetImagePath().GetValue()
	}

	return &Meal{
		Id:                   id,
		Author:               protoEntity.GetAuthor(),
		ImagePath:            imagePath,
		BaseRecipeId:         baseRecipeId,
		WeightedIngredients:  weightedIngredients,
		ConsumptionTimestamp: protoEntity.GetConsumptionTimestamp().GetSeconds(),
	}, nil
}

func MealToProto(mongoEntity *Meal) (*pbNutrition.Meal, error) {
	if mongoEntity == nil {
		return nil, ErrNilInstance
	}

	weightedIngredients := make([]*pbNutrition.WeightedIngredient, 0, len(mongoEntity.WeightedIngredients))
	for _, weightedIngredientMongo := range mongoEntity.WeightedIngredients {
		weightedIngredient, err := WeightedIngredientToProto(weightedIngredientMongo)
		if err != nil {
			return nil, err
		}
		weightedIngredients = append(weightedIngredients, weightedIngredient)
	}

	var baseRecipeId *wrapperspb.StringValue
	if mongoEntity.BaseRecipeId.IsZero() {
		baseRecipeId = &wrapperspb.StringValue{
			Value: mongoEntity.BaseRecipeId.Hex(),
		}
	}

	var imagePath *wrapperspb.StringValue
	if mongoEntity.ImagePath != "" {
		imagePath = &wrapperspb.StringValue{
			Value: mongoEntity.ImagePath,
		}
	}

	return &pbNutrition.Meal{
		Id:                  mongoEntity.Id.Hex(),
		Author:              mongoEntity.Author,
		ImagePath:           imagePath,
		BaseRecipeId:        baseRecipeId,
		WeightedIngredients: weightedIngredients,
		ConsumptionTimestamp: &timestamppb.Timestamp{
			Seconds: mongoEntity.ConsumptionTimestamp,
		},
	}, nil
}
