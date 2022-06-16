package transportation

import (
	"strings"
	"time"
)

const (
	RFC3339     = "2006-01-02T15:04:05"
)

type OcpiTime time.Time

func (t OcpiTime) MarshalJSON() ([]byte, error) {
	ot :=time.Time(t)
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = ot.AppendFormat(b, time.RFC3339)
	b = append(b, '"')

	return b, nil
}

func (t OcpiTime) String() string {
	return time.Time(t).String()
}

func NewOcpiTime(t *time.Time) OcpiTime {
	if t == nil {
		return (OcpiTime)(time.Now())
	}

	return (OcpiTime)(*t)
}

func (t *OcpiTime) UnmarshalJSON(data []byte) error {
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