package versiondetail

import (
	"encoding/json"
	"io"
	"time"
)

type OCPIVersionDetailResponse struct {
	Data          *VersionDetailPayload `json:"data,omitempty"`
	StatusCode    int16                 `json:"status_code"`
	StatusMessage string                `json:"status_message"`
	Timestamp     time.Time             `json:"timestamp"`
}

func (r *VersionDetailResolver) UnmarshalPayload(body io.ReadCloser) (*VersionDetailPayload, error) {
	payload := VersionDetailPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func (r *VersionDetailResolver) UnmarshalResponse(body io.ReadCloser) (*OCPIVersionDetailResponse, error) {
	response := OCPIVersionDetailResponse{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
