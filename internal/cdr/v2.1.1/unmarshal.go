package cdr

import (
	"encoding/json"
	"io"
)

func (r *CdrResolver) UnmarshalDto(body io.ReadCloser) (*CdrDto, error) {
	dto := CdrDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
