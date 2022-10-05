package version

import (
	"encoding/json"
	"io"

	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *VersionResolver) UnmarshalPushDto(body io.ReadCloser) (*coreDto.VersionDto, error) {
	versionDto := coreDto.VersionDto{}

	if err := json.NewDecoder(body).Decode(&versionDto); err != nil {
		return nil, err
	}

	return &versionDto, nil
}

func (r *VersionResolver) UnmarshalPullDto(body io.ReadCloser) (*coreDto.OcpiVersionsDto, error) {
	response := coreDto.OcpiVersionsDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
