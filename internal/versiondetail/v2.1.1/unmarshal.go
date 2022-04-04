package versiondetail

import (
	"encoding/json"
	"io"
	"time"
)

type OCPIVersionDetailResponse struct {
	Data          *VersionDetailDto `json:"data,omitempty"`
	StatusCode    int16             `json:"status_code"`
	StatusMessage string            `json:"status_message"`
	Timestamp     time.Time         `json:"timestamp"`
}

func (r *VersionDetailResolver) UnmarshalDto(body io.ReadCloser) (*VersionDetailDto, error) {
	dto := VersionDetailDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *VersionDetailResolver) UnmarshalResponse(body io.ReadCloser) (*OCPIVersionDetailResponse, error) {
	response := OCPIVersionDetailResponse{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
