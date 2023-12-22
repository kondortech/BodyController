package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/durationpb"
)

type RecipeMongoDB struct {
	Id                      primitive.ObjectID          `bson:"_id,omitempty"`
	Title                   string                      `bson:"title"`
	TasteDescription        string                      `bson:"taste_description"`
	CookingStepsDescription string                      `bson:"cooking_steps_description"`
	Author                  string                      `bson:"author"`
	RequiredIngredients     []WeightedIngredientMongoDB `bson:"required_ingredients"`
	CookingTime             int64                       `bson:"cooking_time"`
	Macros                  MacrosMongoDB               `bson:"macros"`
}

func (proto *Recipe) ConvertToMongoDocument() (*RecipeMongoDB, error) {
	mongoMacros, _ := proto.Macros.ConvertToMongoDocument()

	mongo := &RecipeMongoDB{
		Title:                   proto.Title,
		TasteDescription:        proto.TasteDescription,
		CookingStepsDescription: proto.CookingStepsDescription,
		Author:                  proto.Author,
		RequiredIngredients:     make([]WeightedIngredientMongoDB, 0, len(proto.RequiredIngredients)),
		CookingTime:             proto.CookingTime.Seconds,
		Macros:                  *mongoMacros,
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
	protoMacros, _ := mongo.Macros.ConvertToProtoMessage()
	proto := &Recipe{
		HexId:                   mongo.Id.Hex(),
		Title:                   mongo.Title,
		TasteDescription:        mongo.TasteDescription,
		CookingStepsDescription: mongo.CookingStepsDescription,
		Author:                  mongo.Author,
		RequiredIngredients:     make([]*WeightedIngredient, 0, len(mongo.RequiredIngredients)),
		CookingTime:             &durationpb.Duration{Seconds: mongo.CookingTime},
		Macros:                  protoMacros,
	}

	for _, mongoWeightedIngredient := range mongo.RequiredIngredients {
		protoWeightedIngredient, _ := mongoWeightedIngredient.ConvertToProtoMessage()
		proto.RequiredIngredients = append(proto.RequiredIngredients, protoWeightedIngredient)
	}

	return proto, nil
}
