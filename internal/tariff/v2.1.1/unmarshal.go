package tariff

import (
	"encoding/json"
	"io"
)

func (r *TariffResolver) UnmarshalPushDto(body io.ReadCloser) (*TariffDto, error) {
	dto := TariffDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *TariffResolver) UnmarshalPullDto(body io.ReadCloser) (*OCPITariffsDto, error) {
	dto := OCPITariffsDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
