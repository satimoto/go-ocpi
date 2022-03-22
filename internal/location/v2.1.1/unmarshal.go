package location

import (
	"encoding/json"
	"io"
)

func (r *LocationResolver) UnmarshalPayload(body io.ReadCloser) (*LocationPayload, error) {
	payload := LocationPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
