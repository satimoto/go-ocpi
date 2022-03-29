package session

import (
	"encoding/json"
	"io"
)

func (r *SessionResolver) UnmarshalPayload(body io.ReadCloser) (*SessionPayload, error) {
	payload := SessionPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
