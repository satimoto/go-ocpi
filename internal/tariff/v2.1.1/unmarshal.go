package tariff

import (
	"encoding/json"
	"io"
)

func (r *TariffResolver) UnmarshalPayload(body io.ReadCloser) (*TariffPayload, error) {
	payload := TariffPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
