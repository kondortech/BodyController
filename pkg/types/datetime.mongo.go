package types

type DateTimeMongo struct {
	Year    int32 `bson:"year"`
	Month   int32 `bson:"month"`
	Day     int32 `bson:"day"`
	Hours   int32 `bson:"hours"`
	Minutes int32 `bson:"minutes"`
	Seconds int32 `bson:"seconds"`
}

func (proto *DateTime) ConvertToMongoDocument() (*DateTimeMongo, error) {
	return &DateTimeMongo{
		Year:    proto.Year,
		Month:   proto.Month,
		Day:     proto.Day,
		Hours:   proto.Hours,
		Minutes: proto.Minutes,
		Seconds: proto.Seconds,
	}, nil
}

func (mongo *DateTimeMongo) ConvertToProtoMessage() (*DateTime, error) {
	return &DateTime{
		Year:    mongo.Year,
		Month:   mongo.Month,
		Day:     mongo.Day,
		Hours:   mongo.Hours,
		Minutes: mongo.Minutes,
		Seconds: mongo.Seconds,
	}, nil
}
