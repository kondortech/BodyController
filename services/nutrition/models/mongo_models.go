package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrNilInstance = errors.New("nil instance")
var ErrInvalidId = errors.New("invalid id")

type MacrosMongo struct {
	Proteins float32 `bson:"proteins"`
	Carbs    float32 `bson:"carbs"`
	Fats     float32 `bson:"fats"`
	Calories float32 `bson:"calories"`
}

func (instance *Macros) Mongo() (*MacrosMongo, error) {
	return &MacrosMongo{
		Proteins: float32(instance.GetProteins()),
		Carbs:    float32(instance.GetCarbs()),
		Fats:     float32(instance.GetFats()),
		Calories: float32(instance.GetCalories()),
	}, nil
}

func (instance *MacrosMongo) Proto() (*Macros, error) {
	if instance == nil {
		return nil, ErrNilInstance
	}
	return &Macros{
		Calories: instance.Calories,
		Proteins: instance.Proteins,
		Carbs:    instance.Carbs,
		Fats:     instance.Fats,
	}, nil
}

type MacrosForWeightMongo struct {
	Macros *MacrosMongo `bson:"macros"`
	Gramms float32      `bson:"gramms"`
}

func (instance *MacrosForWeight) Mongo() (*MacrosForWeightMongo, error) {
	macros, err := instance.GetMacros().Mongo()
	if err != nil {
		return nil, err
	}

	return &MacrosForWeightMongo{
		Macros: macros,
		Gramms: instance.GetGramms(),
	}, nil
}

func (instance *MacrosForWeightMongo) Proto() (*MacrosForWeight, error) {
	if instance == nil {
		return nil, ErrNilInstance
	}
	macros, err := instance.Macros.Proto()
	if err != nil {
		return nil, err
	}

	return &MacrosForWeight{
		Macros: macros,
		Gramms: instance.Gramms,
	}, nil
}

type IngredientMongo struct {
	Id              primitive.ObjectID    `bson:"_id,omitempty"`
	Name            string                `bson:"name"`
	MacrosForWeight *MacrosForWeightMongo `bson:"macros_for_weight"`
}

func (instance *Ingredient) Mongo() (*IngredientMongo, error) {
	macrosForWeight, err := instance.GetMacrosForWeight().Mongo()
	if err != nil {
		return nil, err
	}

	id, err := primitive.ObjectIDFromHex(instance.GetId())
	if err != nil {
		return nil, ErrInvalidId
	}

	return &IngredientMongo{
		Id:              id,
		Name:            instance.GetName(),
		MacrosForWeight: macrosForWeight,
	}, nil
}

func (instance *IngredientMongo) Proto() (*Ingredient, error) {
	if instance == nil {
		return nil, ErrNilInstance
	}

	macrosForWeight, err := instance.MacrosForWeight.Proto()
	if err != nil {
		return nil, err
	}

	return &Ingredient{
		Id:              instance.Id.Hex(),
		Name:            instance.Name,
		MacrosForWeight: macrosForWeight,
	}, nil
}

type WeightedIngredientMongo struct {
	Ingredient *IngredientMongo `bson:"ingredient"`
	Gramms     float32          `bson:"gramms"`
}

func (instance *WeightedIngredient) Mongo() (*WeightedIngredientMongo, error) {
	ingredient, err := instance.GetIngredient().Mongo()
	if err != nil {
		return nil, err
	}

	return &WeightedIngredientMongo{
		Ingredient: ingredient,
		Gramms:     instance.GetGramms(),
	}, nil
}

func (instance *WeightedIngredientMongo) Proto() (*WeightedIngredient, error) {
	if instance == nil {
		return nil, ErrNilInstance
	}

	ingredient, err := instance.Ingredient.Proto()
	if err != nil {
		return nil, err
	}

	return &WeightedIngredient{
		Ingredient: ingredient,
		Gramms:     instance.Gramms,
	}, nil
}

type RecipeMongo struct {
	Id                            primitive.ObjectID         `bson:"_id,omitempty"`
	Name                          string                     `bson:"name"`
	RecipeSteps                   string                     `bson:"recipe_steps"`
	ExampleIngredientsProportions []*WeightedIngredientMongo `bson:"example_ingredients_proportions"`
}

func (instance *Recipe) Mongo() (*RecipeMongo, error) {
	id, err := primitive.ObjectIDFromHex(instance.GetId())
	if err != nil {
		return nil, ErrInvalidId
	}

	ingredients := make([]*WeightedIngredientMongo, 0, len(instance.GetExampleIngredientsProportions()))
	for _, weightedIngredient := range instance.GetExampleIngredientsProportions() {
		weightedIngredientMongo, err := weightedIngredient.Mongo()
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, weightedIngredientMongo)
	}

	return &RecipeMongo{
		Id:                            id,
		Name:                          instance.GetName(),
		RecipeSteps:                   instance.GetRecipeSteps(),
		ExampleIngredientsProportions: ingredients,
	}, nil
}

func (instance *RecipeMongo) Proto() (*Recipe, error) {
	if instance == nil {
		return nil, ErrNilInstance
	}

	ingredients := make([]*WeightedIngredient, 0, len(instance.ExampleIngredientsProportions))
	for _, weightedIngredientMongo := range instance.ExampleIngredientsProportions {
		weightedIngredient, err := weightedIngredientMongo.Proto()
		if err != nil {
			return nil, err
		}
		ingredients = append(ingredients, weightedIngredient)
	}

	return &Recipe{
		Id:                            instance.Id.Hex(),
		Name:                          instance.Name,
		RecipeSteps:                   instance.RecipeSteps,
		ExampleIngredientsProportions: ingredients,
	}, nil
}
