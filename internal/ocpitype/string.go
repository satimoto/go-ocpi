package ocpitype

import (
	"strings"
)

type String string

func NewString(str string) String {
	return (String)(str)
}

func (s String) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(s)+2)
	b = append(b, '"')
	b = append(b, []byte(s)...)
	b = append(b, '"')

	return b, nil
}

func (s String) String() string {
	return (string)(s)
}

func (s *String) UnmarshalJSON(data []byte) error {
	dataStr := string(data)

	if dataStr == "null" {
		return nil
	}

	*(*string)(s) = strings.Replace(dataStr, `"`, ``, -1)
	return nil
}