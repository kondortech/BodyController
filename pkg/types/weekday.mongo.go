package types

func EncodeWeekday(weekday Weekday) int32 {
	return int32(weekday)
}

func DecodeWeekday(value int32) Weekday {
	return Weekday(value)
}
