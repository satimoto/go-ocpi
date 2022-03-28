package token

import (
	"encoding/json"
	"io"
)

func (r *TokenResolver) UnmarshalPayload(body io.ReadCloser) (*TokenPayload, error) {
	payload := TokenPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func (r *TokenResolver) UnmarshalLocationReferencesPayload(body io.ReadCloser) (*LocationReferencesPayload, error) {
	if body != nil {
		payload := LocationReferencesPayload{}

		if err := json.NewDecoder(body).Decode(&payload); err != nil {
			return nil, err
		}

		return &payload, nil
	}

	return nil, nil
}
