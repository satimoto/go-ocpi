package transportation

import (
	"encoding/json"
	"io"
)

func UnmarshalResponseDto(body io.ReadCloser) (*OcpiResponse, error) {
	dto := OcpiResponse{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
