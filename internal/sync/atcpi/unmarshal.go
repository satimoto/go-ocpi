package atcpi

import (
	"encoding/json"
	"io"
)

func UnmarshalDto(body io.ReadCloser) ([]*PriceInformationDto, error) {
	response := []*PriceInformationDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
