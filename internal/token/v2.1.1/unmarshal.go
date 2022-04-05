package token

import (
	"encoding/json"
	"io"
)

func (r *TokenResolver) UnmarshalDto(body io.ReadCloser) (*TokenDto, error) {
	dto := TokenDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}
