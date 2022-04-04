package session

import (
	"encoding/json"
	"io"
)

func (r *SessionResolver) UnmarshalDto(body io.ReadCloser) (*SessionDto, error) {
	dto := SessionDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
