package evse

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *EvseResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiEvseDto, error) {
	response := dto.OcpiEvseDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
