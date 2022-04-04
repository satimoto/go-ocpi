package tariff

import (
	"encoding/json"
	"io"
)

func (r *TariffResolver) UnmarshalTariffPullDto(body io.ReadCloser) (*TariffPullDto, error) {
	dto := TariffPullDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *TariffResolver) UnmarshalTariffPushDto(body io.ReadCloser) (*TariffPushDto, error) {
	dto := TariffPushDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
