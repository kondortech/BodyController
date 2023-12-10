package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/durationpb"
)

type MacrosMongoDB struct {
	Proteins float32 `bson:"proteins"`
	Carbs    float32 `bson:"carbs"`
	Fats     float32 `bson:"fats"`
	Calories float32 `bson:"calories"`
}

func (protoMacros *Macros) ConvertToMongoDocument() (*MacrosMongoDB, error) {
	return &MacrosMongoDB{
		Proteins: protoMacros.Proteins,
		Carbs:    protoMacros.Carbs,
		Fats:     protoMacros.Fats,
		Calories: protoMacros.Calories,
	}, nil
}

func (mongoMacros *MacrosMongoDB) ConvertToProtoMessage() (*Macros, error) {
	return &Macros{
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

type RecipeMongoDB struct {
	Id                      primitive.ObjectID `bson:"_id,omitempty"`
	Title                   string             `bson:"title"`
	TasteDescription        string             `bson:"taste_description"`
	CookingStepsDescription string             `bson:"cooking_steps_description"`
	Author                  string             `bson:"author"`

	RequiredIngredients []WeightedIngredientMongoDB `protobuf:"bytes,6,rep,name=required_ingredients,json=requiredIngredients,proto3" json:"required_ingredients,omitempty"`
	CookingTime         int64
}

func (proto *Recipe) ConvertToMongoDocument() (*RecipeMongoDB, error) {
	mongo := &RecipeMongoDB{
		Title:                   proto.Title,
		TasteDescription:        proto.TasteDescription,
		CookingStepsDescription: proto.CookingStepsDescription,
		Author:                  proto.Author,
		RequiredIngredients:     make([]WeightedIngredientMongoDB, 0, len(proto.RequiredIngredients)),
		CookingTime:             proto.CookingTime.Seconds,
	}

	for _, protoWeightedIngredient := range proto.RequiredIngredients {
		mongoWeightedIngredient, _ := protoWeightedIngredient.ConvertToMongoDocument()
		mongo.RequiredIngredients = append(mongo.RequiredIngredients, *mongoWeightedIngredient)
	}

	if len(proto.GetHexId()) != 0 {
		objectId, err := primitive.ObjectIDFromHex(proto.GetHexId())
		if err != nil {
			return nil, fmt.Errorf("Recipe.ConvertToMongoDocument returned error: %v", err)
		}
		mongo.Id = objectId
	}

	return mongo, nil
}

func (mongo *RecipeMongoDB) ConvertToProtoMessage() (*Recipe, error) {
	proto := &Recipe{
		HexId:                   mongo.Id.Hex(),
		Title:                   mongo.Title,
		TasteDescription:        mongo.TasteDescription,
		CookingStepsDescription: mongo.CookingStepsDescription,
		Author:                  mongo.Author,
		RequiredIngredients:     make([]*WeightedIngredient, 0, len(mongo.RequiredIngredients)),
		CookingTime:             &durationpb.Duration{Seconds: mongo.CookingTime},
	}

	for _, mongoWeightedIngredient := range mongo.RequiredIngredients {
		protoWeightedIngredient, _ := mongoWeightedIngredient.ConvertToProtoMessage()
		proto.RequiredIngredients = append(proto.RequiredIngredients, protoWeightedIngredient)
	}

	return proto, nil
}
