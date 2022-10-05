package tokenauthorization

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *TokenAuthorizationResolver) UnmarshalLocationReferencesDto(body io.ReadCloser) (*dto.LocationReferencesDto, error) {
	if body != nil {
		locationReferencesDto := dto.LocationReferencesDto{}

		if err := json.NewDecoder(body).Decode(&locationReferencesDto); err != nil {
			return nil, err
		}

		return &locationReferencesDto, nil
	}

	return nil, nil
}
