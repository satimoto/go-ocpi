package util

import (
	"database/sql"
	"time"
)

func NilBool(i interface{}) *bool {
	switch t := i.(type) {
	case sql.NullBool:
		if t.Valid {
			return &t.Bool
		}
	case bool:
		return &t
	case *bool:
		return t
	}

	return nil
}

func NilFloat64(i interface{}) *float64 {
	switch t := i.(type) {
	case sql.NullFloat64:
		if t.Valid {
			return &t.Float64
		}
	case float64:
		if t != 0 {
			return &t
		}
	case *float64:
		return t
	}

	return nil
}

func NilInt16(i int16) *int16 {
	if i == 0 {
		return nil
	}
	return &i
}

func NilInt32(i int32) *int32 {
	if i == 0 {
		return nil
	}
	return &i
}

func NilInt64(i interface{}) *int64 {
	switch t := i.(type) {
	case sql.NullInt64:
		if t.Valid {
			return &t.Int64
		}
	case int64:
		if t != 0 {
			return &t
		}
	case *int64:
		return t
	}

	return nil
}

func NilString(i interface{}) *string {
	switch t := i.(type) {
	case sql.NullString:
		if t.Valid {
			return &t.String
		}
	case string:
		return &t
	case *string:
		return t
	}

	return nil
}

func NilTime(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}
