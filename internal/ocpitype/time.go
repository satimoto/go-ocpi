package ocpitype

import (
	"strings"
	"time"

	"github.com/satimoto/go-datastore/pkg/util"
)

const (
	RFC3339     = "2006-01-02T15:04:05Z"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	ot :=time.Time(t).UTC()
	b := make([]byte, 0, len(RFC3339)+2)
	b = append(b, '"')
	b = ot.AppendFormat(b, RFC3339)
	b = append(b, '"')

	return b, nil
}

func (t Time) String() string {
	return time.Time(t).String()
}

func (t Time) Time() time.Time {
	return (time.Time)(t)
}

func ParseOcpiTime(str string, fallback *time.Time) *Time {
	t := util.ParseTime(str, fallback)

	return NilOcpiTime(t)
}

func NewOcpiTime(t *time.Time) Time {
	ot := NilOcpiTime(t)

	if ot == nil {
		return (Time)(time.Now())
	}
	
	return *ot
}

func NilOcpiTime(t *time.Time) *Time {
	if t != nil {
		if t.IsZero() {
			return nil
		}
		
		ot := (Time)(*t)
		return &ot
	}

	return nil
}

func NilTime(t *Time) *time.Time {
	if t != nil {
		return util.NilTime(t.Time())
	}

	return nil
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