package models

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
