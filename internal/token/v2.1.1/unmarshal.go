package token

import (
	"encoding/json"
	"io"

	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *TokenResolver) UnmarshalDto(body io.ReadCloser) (*TokenDto, error) {
	dto := TokenDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *TokenResolver) UnmarshalPullDto(body io.ReadCloser) (*transportation.OcpiResponse, error) {
	dto := transportation.OcpiResponse{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
