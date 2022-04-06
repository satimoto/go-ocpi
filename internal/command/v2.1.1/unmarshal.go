package command

import (
	"encoding/json"
	"io"
)

func (r *CommandResolver) UnmarshalPushDto(body io.ReadCloser) (*CommandResponseDto, error) {
	if body != nil {
		dto := CommandResponseDto{}

		if err := json.NewDecoder(body).Decode(&dto); err != nil {
			return nil, err
		}

		return &dto, nil
	}

	return nil, nil
}
