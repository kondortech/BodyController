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

func EncodeMeasureUnit(unit MeasureUnit) int32 {
	return int32(unit)
}

func DecodeMeasureUnit(unit int32) MeasureUnit {
	return MeasureUnit(unit)
}

type QuantityMongoDB struct {
	MeasureUnit int32   `bson:"measure_unit"`
	Quantity    float32 `bson:"quantity"`
}

func (proto *Quantity) ConvertToMongoDocument() (*QuantityMongoDB, error) {
	return &QuantityMongoDB{
		Quantity:    proto.Quantity,
		MeasureUnit: EncodeMeasureUnit(proto.MeasureUnit),
	}, nil
}

func (mongo *QuantityMongoDB) ConvertToProtoMessage() (*Quantity, error) {
	return &Quantity{
		MeasureUnit: DecodeMeasureUnit(mongo.MeasureUnit),
		Quantity:    mongo.Quantity,
	}, nil
}

type WeightedIngredientMongoDB struct {
	IngredientHexId primitive.ObjectID `bson:"ingredient_hex_id"`
	Quantity        QuantityMongoDB    `bson:"quantity"`
}

func (proto *WeightedIngredient) ConvertToMongoDocument() (*WeightedIngredientMongoDB, error) {
	mongoIngredient, _ := proto.Ingredient.ConvertToMongoDocument()
	mongoQuantity, _ := proto.Quantity.ConvertToMongoDocument()

	return &WeightedIngredientMongoDB{
		IngredientHexId: mongoIngredient.Id,
		Quantity:        *mongoQuantity,
	}, nil
}

func (mongo *WeightedIngredientMongoDB) ConvertToProtoMessage() (*WeightedIngredient, error) {
	protoIngredient := &Ingredient{
		HexId: mongo.IngredientHexId.Hex(),
	}
	protoQuantity, _ := mongo.Quantity.ConvertToProtoMessage()
	return &WeightedIngredient{
		Ingredient: protoIngredient,
		Quantity:   protoQuantity,
	}, nil
}
