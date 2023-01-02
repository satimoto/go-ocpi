package cdr

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *CdrResolver) UnmarshalPushDto(body io.ReadCloser) (*dto.CdrDto, error) {
	cdrDto := dto.CdrDto{}

	if err := json.NewDecoder(body).Decode(&cdrDto); err != nil {
		return nil, err
	}

	return &cdrDto, nil
}

func (r *CdrResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiCdrsDto, error) {
	response := dto.OcpiCdrsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
