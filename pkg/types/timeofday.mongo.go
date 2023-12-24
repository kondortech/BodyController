package types

type TimeOfDayMongo struct {
	Hours   int32 `bson:"hours"`
	Minutes int32 `bson:"minutes"`
	Seconds int32 `bson:"seconds"`
}

func (proto *TimeOfDay) ConvertToMongoDocument() (*TimeOfDayMongo, error) {
	return &TimeOfDayMongo{
		Hours:   proto.Hours,
		Minutes: proto.Minutes,
		Seconds: proto.Seconds,
	}, nil
}

func (mongo *TimeOfDayMongo) ConvertToProtoMessage() (*TimeOfDay, error) {
	return &TimeOfDay{
		Hours:   mongo.Hours,
		Minutes: mongo.Minutes,
		Seconds: mongo.Seconds,
	}, nil
}
