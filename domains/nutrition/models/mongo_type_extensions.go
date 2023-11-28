package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MacrosMongoDB struct {
	Proteins float32 `bson:"proteins"`
	Carbs    float32 `bson:"carbs"`
	Fats     float32 `bson:"fats"`
	Calories float32 `bson:"calories"`
}

func (protoMacros *Macros100G) ConvertToMongoReadable() (*MacrosMongoDB, error) {
	return &MacrosMongoDB{
		Proteins: protoMacros.Proteins,
		Carbs:    protoMacros.Carbs,
		Fats:     protoMacros.Fats,
		Calories: protoMacros.Calories,
	}, nil
}

func (mongoMacros *MacrosMongoDB) ConvertToProtoMessage() (*Macros100G, error) {
	return &Macros100G{
		Proteins: mongoMacros.Proteins,
		Carbs:    mongoMacros.Carbs,
		Fats:     mongoMacros.Fats,
		Calories: mongoMacros.Calories,
	}, nil
}

type IngredientMongoDB struct {
	Id primitive.ObjectID `bson:"_id,omitempty"`

	MongoMacros MacrosMongoDB `bson:"macros_100g"`
	Title       string        `bson:"title"`
	Description string        `bson:"description"`
	Author      string        `bson:"author"`
}

func (protoIngredient *Ingredient) ConvertToMongoDocument() (*IngredientMongoDB, error) {
	mongoMacros, err := protoIngredient.GetMacros100G().ConvertToMongoReadable()
	if err != nil {
		return nil, fmt.Errorf("Ingredient.ConvertToMongoDocument returned error: %v", err)
	}

	mongoIngredient := &IngredientMongoDB{
		MongoMacros: *mongoMacros,
		Title:       protoIngredient.Title,
		Description: protoIngredient.Description,
		Author:      protoIngredient.Author,
	}
	if len(protoIngredient.GetHexId()) != 0 {
		objectId, err := primitive.ObjectIDFromHex(protoIngredient.GetHexId())
		if err != nil {
			return nil, fmt.Errorf("Ingredient.ConvertToMongoDocument returned error: %v", err)
		}
		mongoIngredient.Id = objectId
	}
	return mongoIngredient, nil
}

func (mongoIngredient *IngredientMongoDB) ConvertToProtoMessage() (*Ingredient, error) {
	protoMacros, err := mongoIngredient.MongoMacros.ConvertToProtoMessage()
	if err != nil {
		return nil, fmt.Errorf("IngredientMongoDB.ConvertToProtoMessage returned error: %v", err)
	}

	return &Ingredient{
		HexId:       mongoIngredient.Id.Hex(),
		Macros100G:  protoMacros,
		Title:       mongoIngredient.Title,
		Description: mongoIngredient.Description,
		Author:      mongoIngredient.Author,
	}, nil
}
