package command

import (
	"encoding/json"
	"io"

	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *CommandResolver) UnmarshalPushDto(body io.ReadCloser) (*dto.CommandResponseDto, error) {
	if body != nil {
		commandResponseDto := dto.CommandResponseDto{}

		if err := json.NewDecoder(body).Decode(&commandResponseDto); err != nil {
			return nil, err
		}

		return &commandResponseDto, nil
	}

	return nil, nil
}

func (r *CommandResolver) UnmarshalPullDto(body io.ReadCloser) (*dto.OcpiCommandResponseDto, error) {
	if body != nil {
		response := dto.OcpiCommandResponseDto{}

		if err := json.NewDecoder(body).Decode(&response); err != nil {
			return nil, err
		}

		return &response, nil
	}

	return nil, nil
}
