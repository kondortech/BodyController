package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func EncodeWeekday(weekday Weekday) int32 {
	return int32(weekday)
}

func DecodeWeekday(value int32) Weekday {
	return Weekday(value)
}

func ConvertToMongoField()

type NutritionStepMongoDB struct {
	MacrosRequirements MacrosMongoDB `bson:"macros_requirements"`
	StartTimestamp     int64         `bson:"start_timestamp"`
	FinishTimestamp    int64         `bson:"finish_timestamp"`
	Title              string        `bson:"title"`
}

func (proto *NutritionStep) ConvertToMongoDocument() (*NutritionStepMongoDB, error) {
	mongoMacros, _ := proto.MacrosRequirements.ConvertToMongoDocument()

	return &NutritionStepMongoDB{
		MacrosRequirements: *mongoMacros,
		StartTimestamp:     proto.StartTimestamp.Seconds,
		FinishTimestamp:    proto.FinishTimestamp.Seconds,
		Title:              proto.Title,
	}, nil
}

func (mongo *NutritionStepMongoDB) ConvertToProtoMessage() (*NutritionStep, error) {
	protoMacros, _ := mongo.MacrosRequirements.ConvertToProtoMessage()

	return &NutritionStep{
		MacrosRequirements: protoMacros,
		StartTimestamp:     &timestamppb.Timestamp{Seconds: mongo.StartTimestamp},
		FinishTimestamp:    &timestamppb.Timestamp{Seconds: mongo.FinishTimestamp},
		Title:              mongo.Title,
	}, nil
}

type NutritionLifestyleTemplateMongoDB struct {
	HexId              primitive.ObjectID     `bson:"_id,omitempty"`
	Title              string                 `bson:"title"`
	Author             string                 `bson:"author"`
	CycleDurationDays  int64                  `bson:"cycle_duration_days"`
	NutritionSteps     []NutritionStepMongoDB `bson:"nutrition_steps"`
	AverageDailyMacros MacrosMongoDB          `bson:"average_daily_macros"`
	StartingWeekday    int32                  `bson:"starting_weekday"`
}

func (proto *NutritionLifestyleTemplate) ConvertToMongoDocument() (*NutritionLifestyleTemplateMongoDB, error) {
	mongoMacros, _ := proto.AverageDailyMacros.ConvertToMongoDocument()

	mongo := &NutritionLifestyleTemplateMongoDB{
		Title:              proto.Title,
		Author:             proto.Title,
		CycleDurationDays:  proto.CycleDurationDays,
		NutritionSteps:     make([]NutritionStepMongoDB, 0, len(proto.NutritionSteps)),
		AverageDailyMacros: *mongoMacros,
		StartingWeekday:    EncodeWeekday(proto.StartingWeekday),
	}

	for _, protoNutritionStep := range proto.NutritionSteps {
		mongoNutritionStep, _ := protoNutritionStep.ConvertToMongoDocument()
		mongo.NutritionSteps = append(mongo.NutritionSteps, *mongoNutritionStep)
	}

	if len(proto.GetHexId()) != 0 {
		objectId, err := primitive.ObjectIDFromHex(proto.GetHexId())
		if err != nil {
			return nil, fmt.Errorf("Recipe.ConvertToMongoDocument returned error: %v", err)
		}
		mongo.HexId = objectId
	}

	return mongo, nil
}

func (mongo *NutritionLifestyleTemplateMongoDB) ConvertToProtoMessage() (*NutritionLifestyleTemplate, error) {
	protoMacros, _ := mongo.AverageDailyMacros.ConvertToProtoMessage()
	proto := &NutritionLifestyleTemplate{
		HexId:              mongo.HexId.Hex(),
		Title:              mongo.Title,
		Author:             mongo.Author,
		CycleDurationDays:  mongo.CycleDurationDays,
		NutritionSteps:     make([]*NutritionStep, 0, len(mongo.NutritionSteps)),
		AverageDailyMacros: protoMacros,
		StartingWeekday:    DecodeWeekday(mongo.StartingWeekday),
	}

	for _, mongoNutritionStep := range mongo.NutritionSteps {
		protoNutritionStep, _ := mongoNutritionStep.ConvertToProtoMessage()
		proto.NutritionSteps = append(proto.NutritionSteps, protoNutritionStep)
	}

	return proto, nil
}
