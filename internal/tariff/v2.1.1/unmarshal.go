package tariff

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *TariffResolver) UnmarshalPushDto(body io.ReadCloser) (*dto.TariffDto, error) {
	dto := dto.TariffDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *TariffResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiTariffsDto, error) {
	response := dto.OcpiTariffsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
