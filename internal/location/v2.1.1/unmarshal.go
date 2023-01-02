package location

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *LocationResolver) UnmarshalPushDto(body io.ReadCloser) (*dto.LocationDto, error) {
	locationDto := dto.LocationDto{}

	if err := json.NewDecoder(body).Decode(&locationDto); err != nil {
		return nil, err
	}

	return &locationDto, nil
}

func (r *LocationResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiLocationsDto, error) {
	response := dto.OcpiLocationsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
