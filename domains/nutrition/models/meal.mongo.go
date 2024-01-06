package models

import (
	"errors"
	"fmt"

	pbTypes "github.com/kirvader/BodyController/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EncodeMealStatus(mealStatus MealStatus) string {
	return mealStatus.String()
}

func DecodeMealStatus(encodedMealString string) (MealStatus, error) {
	if intValue, ok := MealStatus_value[encodedMealString]; ok {
		return MealStatus(intValue), nil
	}
	return MealStatus_NOT_SET, fmt.Errorf("DecodeMealStatus: MealStatus couldn't decode %q", encodedMealString)

}

type MealMongo struct {
	Username         string                 `bson:"username"`
	MealDate         *pbTypes.DateMongo     `bson:"meal_date"`
	MealSetIndex     int64                  `bson:"meal_set_index"`
	StatusUpdateTime *pbTypes.DateTimeMongo `bson:"status_update_time"`
	MealStatus       string                 `bson:"meal_status"`
	RecipeId         primitive.ObjectID     `bson:"recipe_id"`
	FoodAmountGramms float32                `bson:"food_amount_gramms"`
}

func (proto *Meal) ConvertToMongoDocument() (*MealMongo, error) {
	statusUpdateTime, _ := proto.StatusUpdateTime.ConvertToMongoDocument()

	if len(proto.GetRecipeId()) == 0 {
		return nil, errors.New("Meal.ConvertToMongoDocument: proto.RecipeId is empty")
	}
	recipeId, err := primitive.ObjectIDFromHex(proto.GetRecipeId())
	if err != nil {
		return nil, fmt.Errorf("Meal.ConvertToMongoDocument: returned error when parsing recipe_id: %v", err)
	}

	mealDateMongo, _ := proto.MealDate.ConvertToMongoDocument()

	return &MealMongo{
		Username:         proto.Username,
		MealDate:         mealDateMongo,
		StatusUpdateTime: statusUpdateTime,
		MealStatus:       EncodeMealStatus(proto.MealStatus),
		RecipeId:         recipeId,
		FoodAmountGramms: proto.FoodAmountGramms,
	}, nil
}

func (mongo *MealMongo) ConvertToProtoMessage() (*Meal, error) {
	mealStatus, err := DecodeMealStatus(mongo.MealStatus)
	if err != nil {
		return nil, err
	}
	mealDateProto, _ := mongo.MealDate.ConvertToProtoMessage()
	statusUpdateTimeProto, _ := mongo.StatusUpdateTime.ConvertToProtoMessage()

	return &Meal{
		Username:         mongo.Username,
		MealDate:         mealDateProto,
		StatusUpdateTime: statusUpdateTimeProto,
		MealStatus:       mealStatus,
		RecipeId:         mongo.RecipeId.Hex(),
		FoodAmountGramms: mongo.FoodAmountGramms,
	}, nil
}
