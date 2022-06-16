package geolocation

import (
	"strings"
)

type DecimalString string

func NewDecimalString(str string) DecimalString {
	return (DecimalString)(str)
}

func (ds DecimalString) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(ds)+2)
	b = append(b, '"')
	b = append(b, []byte(ds)...)
	b = append(b, '"')

	return b, nil
}

func (ds DecimalString) String() string {
	return (string)(ds)
}

func (ds *DecimalString) UnmarshalJSON(data []byte) error {
	dataStr := string(data)

	if dataStr == "null" {
		return nil
	}

	*(*string)(ds) = strings.Replace(dataStr, `"`, ``, -1)
	return nil
}