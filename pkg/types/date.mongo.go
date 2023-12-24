package types

type DateMongo struct {
	Year  int32 `bson:"year"`
	Month int32 `bson:"month"`
	Day   int32 `bson:"day"`
}

func (proto *Date) ConvertToMongoDocument() (*DateMongo, error) {
	if proto == nil {
		return nil, nil
	}
	return &DateMongo{
		Year:  proto.Year,
		Month: proto.Month,
		Day:   proto.Day,
	}, nil
}

func (mongo *DateMongo) ConvertToProtoMessage() (*Date, error) {
	if mongo == nil {
		return nil, nil
	}
	return &Date{
		Year:  mongo.Year,
		Month: mongo.Month,
		Day:   mongo.Day,
	}, nil
}
