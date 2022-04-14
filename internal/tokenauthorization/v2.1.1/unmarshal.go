package tokenauthorization

import (
	"encoding/json"
	"io"
)

func (r *TokenAuthorizationResolver) UnmarshalLocationReferencesDto(body io.ReadCloser) (*LocationReferencesDto, error) {
	if body != nil {
		dto := LocationReferencesDto{}

		if err := json.NewDecoder(body).Decode(&dto); err != nil {
			return nil, err
		}

		return &dto, nil
	}

	return nil, nil
}
