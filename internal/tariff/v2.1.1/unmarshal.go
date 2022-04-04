package tariff

import (
	"encoding/json"
	"io"
)

func (r *TariffResolver) UnmarshalTariffPullPayload(body io.ReadCloser) (*TariffPullPayload, error) {
	payload := TariffPullPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func (r *TariffResolver) UnmarshalTariffPushPayload(body io.ReadCloser) (*TariffPushPayload, error) {
	payload := TariffPushPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
