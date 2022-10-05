package session

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *SessionResolver) UnmarshalPushDto(body io.ReadCloser) (*dto.SessionDto, error) {
	sessionDto := dto.SessionDto{}

	if err := json.NewDecoder(body).Decode(&sessionDto); err != nil {
		return nil, err
	}

	return &sessionDto, nil
}

func (r *SessionResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiSessionsDto, error) {
	response := dto.OcpiSessionsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
