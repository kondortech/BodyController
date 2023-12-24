package models

import (
	"fmt"

	pbTypes "github.com/kirvader/BodyController/pkg/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonalNutritionLifestyleMongoDB struct {
	HexId                           primitive.ObjectID `bson:"_id,omitempty"`
	Title                           string             `bson:"title"`
	Username                        string             `bson:"username"`
	NutritionLifestyleTemplateHexId primitive.ObjectID `bson:"nutrition_lifestyle_template_id"`
	FirstDay                        *pbTypes.DateMongo `bson:"first_day"`
	LastDay                         *pbTypes.DateMongo `bson:"last_day,omitempty"`
}

func (proto *PersonalNutritionLifestyle) ConvertToMongoDocument() (*PersonalNutritionLifestyleMongoDB, error) {
	firstDayMongo, _ := proto.FirstDay.ConvertToMongoDocument()
	lastDayMongo, _ := proto.LastDay.ConvertToMongoDocument()

	mongo := &PersonalNutritionLifestyleMongoDB{
		Title:    proto.Title,
		Username: proto.Username,
		FirstDay: firstDayMongo,
		LastDay:  lastDayMongo,
	}

	if len(proto.GetHexId()) != 0 {
		objectId, err := primitive.ObjectIDFromHex(proto.GetHexId())
		if err != nil {
			return nil, fmt.Errorf("PersonalNutritionLifestyle.ConvertToMongoDocument returned error: %v", err)
		}
		mongo.HexId = objectId
	}

	if len(proto.GetNutritionLifestyleTemplateHexId()) != 0 {
		objectId, err := primitive.ObjectIDFromHex(proto.GetNutritionLifestyleTemplateHexId())
		if err != nil {
			return nil, fmt.Errorf("PersonalNutritionLifestyle.ConvertToMongoDocument returned error: %v", err)
		}
		mongo.NutritionLifestyleTemplateHexId = objectId
	}

	return mongo, nil
}

func (mongo *PersonalNutritionLifestyleMongoDB) ConvertToProtoMessage() (*PersonalNutritionLifestyle, error) {
	firstDayProto, _ := mongo.FirstDay.ConvertToProtoMessage()
	lastDayProto, _ := mongo.LastDay.ConvertToProtoMessage()

	proto := &PersonalNutritionLifestyle{
		HexId:                           mongo.HexId.Hex(),
		Title:                           mongo.Title,
		Username:                        mongo.Username,
		NutritionLifestyleTemplateHexId: mongo.NutritionLifestyleTemplateHexId.Hex(),
		FirstDay:                        firstDayProto,
		LastDay:                         lastDayProto,
	}

	return proto, nil
}
