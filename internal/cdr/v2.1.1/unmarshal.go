package cdr

import (
	"encoding/json"
	"io"
)

func (r *CdrResolver) UnmarshalPushDto(body io.ReadCloser) (*CdrDto, error) {
	dto := CdrDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *CdrResolver) UnmarshalPullDto(body io.ReadCloser) (*OcpiCdrsDto, error) {
	response := OcpiCdrsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
