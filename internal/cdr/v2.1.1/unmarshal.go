package cdr

import (
	"encoding/json"
	"io"
)

func (r *CdrResolver) UnmarshalPayload(body io.ReadCloser) (*CdrPayload, error) {
	payload := CdrPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
