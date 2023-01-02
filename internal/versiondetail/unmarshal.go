package versiondetail

import (
	"encoding/json"
	"io"

	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *VersionDetailResolver) UnmarshalPushDto(body io.ReadCloser) (*coreDto.VersionDetailDto, error) {
	versionDetailDto := coreDto.VersionDetailDto{}

	if err := json.NewDecoder(body).Decode(&versionDetailDto); err != nil {
		return nil, err
	}

	return &versionDetailDto, nil
}

func (r *VersionDetailResolver) UnmarshalPullDto(body io.ReadCloser) (*coreDto.OcpiVersionDetailDto, error) {
	response := coreDto.OcpiVersionDetailDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
