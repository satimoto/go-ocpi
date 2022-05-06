package session

import (
	"encoding/json"
	"io"
)

func (r *SessionResolver) UnmarshalPushDto(body io.ReadCloser) (*SessionDto, error) {
	dto := SessionDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *SessionResolver) UnmarshalPullDto(body io.ReadCloser) (*OcpiSessionsDto, error) {
	response := OcpiSessionsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
