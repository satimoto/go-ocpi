package ocpitype

import (
	"strings"
	"time"
)

const (
	RFC3339     = "2006-01-02T15:04:05"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	ot :=time.Time(t)
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = ot.AppendFormat(b, time.RFC3339)
	b = append(b, '"')

	return b, nil
}

func (t Time) String() string {
	return time.Time(t).String()
}

func NewTime(t *time.Time) Time {
	if t == nil {
		return (Time)(time.Now())
	}

	return (Time)(*t)
}

func (t *Time) UnmarshalJSON(data []byte) error {
	dataStr := string(data)

	if dataStr == "null" {
		return nil
	}

	dataStr = strings.Replace(dataStr, `"`, ``, -1)
	parsedTime, err := time.Parse(time.RFC3339, dataStr)

	if err != nil {
		parsedTime, err = time.Parse(RFC3339, dataStr)

		if err != nil {
			return err
		}
	}

	*(*time.Time)(t) = parsedTime
	return nil
}