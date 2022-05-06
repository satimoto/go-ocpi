package location

import (
	"encoding/json"
	"io"
)

func (r *LocationResolver) UnmarshalPushDto(body io.ReadCloser) (*LocationDto, error) {
	dto := LocationDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *LocationResolver) UnmarshalPullDto(body io.ReadCloser) (*OcpiLocationsDto, error) {
	response := OcpiLocationsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
