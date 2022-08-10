package versiondetail

import (
	"encoding/json"
	"io"
)

func (r *VersionDetailResolver) UnmarshalPushDto(body io.ReadCloser) (*VersionDetailDto, error) {
	dto := VersionDetailDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *VersionDetailResolver) UnmarshalPullDto(body io.ReadCloser) (*OcpiVersionDetailDto, error) {
	response := OcpiVersionDetailDto{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
