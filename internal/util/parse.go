package util

import (
	"strconv"
	"time"
)

func ParseFloat64(str string, defVal float64) float64 {
	f, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return defVal
	}

	return f
}

func ParseInt32(str string, defVal int32) int32 {
	f, err := strconv.ParseInt(str, 10, 32)

	if err != nil {
		return defVal
	}

	return int32(f)
}

func ParseInt64(str string, defVal int64) int64 {
	f, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return defVal
	}

	return f
}

func ParseTime(str string) *time.Time {
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		return nil
	}

	return &t
}