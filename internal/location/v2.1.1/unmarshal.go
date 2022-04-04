package location

import (
	"encoding/json"
	"io"
)

func (r *LocationResolver) UnmarshalDto(body io.ReadCloser) (*LocationDto, error) {
	dto := LocationDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
