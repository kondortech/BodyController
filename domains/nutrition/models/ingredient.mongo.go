package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientMongoDB struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	MongoMacros MacrosMongoDB      `bson:"macros_100g"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Author      string             `bson:"author"`
}

func (protoIngredient *Ingredient) ConvertToMongoDocument() (*IngredientMongoDB, error) {
	mongoMacros, err := protoIngredient.GetMacros100G().ConvertToMongoDocument()
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

type WeightedIngredientMongoDB struct {
	IngredientHexId primitive.ObjectID `bson:"ingredient_hex_id"`
	AmountGramms    float32            `bson:"amount_gramms"`
}

func (proto *WeightedIngredient) ConvertToMongoDocument() (*WeightedIngredientMongoDB, error) {
	mongoIngredient, _ := proto.Ingredient.ConvertToMongoDocument()

	return &WeightedIngredientMongoDB{
		IngredientHexId: mongoIngredient.Id,
		AmountGramms:    proto.AmountGramms,
	}, nil
}

func (mongo *WeightedIngredientMongoDB) ConvertToProtoMessage() (*WeightedIngredient, error) {
	protoIngredient := &Ingredient{
		HexId: mongo.IngredientHexId.Hex(),
	}
	return &WeightedIngredient{
		Ingredient:   protoIngredient,
		AmountGramms: mongo.AmountGramms,
	}, nil
}
