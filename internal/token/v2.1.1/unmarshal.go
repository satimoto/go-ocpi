package token

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TokenResolver) UnmarshalDto(body io.ReadCloser) (*dto.TokenDto, error) {
	dto := dto.TokenDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *TokenResolver) UnmarshalPullDto(body io.ReadCloser) (*transportation.OcpiResponse, error) {
	response := transportation.OcpiResponse{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
