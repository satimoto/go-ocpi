package ocpitype

import (
	"strconv"
	"strings"
)

type Bool bool

func NewBool(b bool) Bool {
	return (Bool)(b)
}

func (b Bool) MarshalJSON() ([]byte, error) {
	if b {
		return []byte("true"), nil
	}

	return []byte("false"), nil
}

func (b Bool) Bool() bool {
	return (bool)(b)
}

func (b *Bool) UnmarshalJSON(data []byte) error {
	dataStr := strings.Replace(string(data), `"`, ``, -1)

	if dataStr == "null" {
		return nil
	}

	dataBool, err := strconv.ParseBool(dataStr)

	if err == nil {
		*(*bool)(b) = dataBool
	}

	return nil
}