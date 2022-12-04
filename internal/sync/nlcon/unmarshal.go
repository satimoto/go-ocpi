package nlcon

import (
	"encoding/json"
	"io"
)

func UnmarshalDto(body io.ReadCloser) ([]*NlConTariffDto, error) {
	response := []*NlConTariffDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
