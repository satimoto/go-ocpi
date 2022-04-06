package version

import (
	"encoding/json"
	"io"
)


func (r *VersionResolver) UnmarshalPushDto(body io.ReadCloser) (*VersionDto, error) {
	dto := VersionDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *VersionResolver) UnmarshalPullDto(body io.ReadCloser) (*OCPIVersionsDto, error) {
	response := OCPIVersionsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
