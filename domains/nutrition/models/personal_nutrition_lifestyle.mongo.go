package models

import (
	"fmt"

	pbTypes "github.com/kirvader/BodyController/pkg/types"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PersonalNutritionLifestyleMongoDB struct {
	HexId                        primitive.ObjectID `bson:"_id,omitempty"`
	Title                        string             `bson:"title"`
	Username                     string             `bson:"username"`
	FirstDay                     *pbTypes.DateMongo `bson:"first_day"`
	LastDay                      *pbTypes.DateMongo `bson:"last_day,omitempty"`
	LowerboundMacrosRequirements *MacrosMongoDB     `bson:"lowerbound_macros_requirements,omitempty"`
	UpperboundMacrosRequirements *MacrosMongoDB     `bson:"upperbound_macros_requirements,omitempty"`
}

func (proto *PersonalNutritionLifestyle) ConvertToMongoDocument() (*PersonalNutritionLifestyleMongoDB, error) {
	firstDayMongo, _ := proto.FirstDay.ConvertToMongoDocument()
	lastDayMongo, _ := proto.LastDay.ConvertToMongoDocument()

	lowerboundMacrosMongo, _ := proto.LowerboundMacrosRequirements.ConvertToMongoDocument()
	upperboundMacrosMongo, _ := proto.UpperboundMacrosRequirements.ConvertToMongoDocument()

	mongo := &PersonalNutritionLifestyleMongoDB{
		Title:                        proto.Title,
		Username:                     proto.Username,
		FirstDay:                     firstDayMongo,
		LastDay:                      lastDayMongo,
		LowerboundMacrosRequirements: lowerboundMacrosMongo,
		UpperboundMacrosRequirements: upperboundMacrosMongo,
	}

	if len(proto.GetHexId()) != 0 {
		objectId, err := primitive.ObjectIDFromHex(proto.GetHexId())
		if err != nil {
			return nil, fmt.Errorf("PersonalNutritionLifestyle.ConvertToMongoDocument returned error: %v", err)
		}
		mongo.HexId = objectId
	}

	return mongo, nil
}

func (mongo *PersonalNutritionLifestyleMongoDB) ConvertToProtoMessage() (*PersonalNutritionLifestyle, error) {
	firstDayProto, _ := mongo.FirstDay.ConvertToProtoMessage()
	lastDayProto, _ := mongo.LastDay.ConvertToProtoMessage()

	lowerboundMacrosProto, _ := mongo.LowerboundMacrosRequirements.ConvertToProtoMessage()
	upperboundMacrosProto, _ := mongo.UpperboundMacrosRequirements.ConvertToProtoMessage()

	proto := &PersonalNutritionLifestyle{
		HexId:                        mongo.HexId.Hex(),
		Title:                        mongo.Title,
		Username:                     mongo.Username,
		FirstDay:                     firstDayProto,
		LastDay:                      lastDayProto,
		LowerboundMacrosRequirements: lowerboundMacrosProto,
		UpperboundMacrosRequirements: upperboundMacrosProto,
	}

	return proto, nil
}
