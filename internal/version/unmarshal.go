package version

import (
	"encoding/json"
	"io"
	"time"
)

type OCPIVersionsResponse struct {
	Data          []*VersionPayload `json:"data,omitempty"`
	StatusCode    int16             `json:"status_code"`
	StatusMessage string            `json:"status_message"`
	Timestamp     time.Time         `json:"timestamp"`
}

func (r *VersionResolver) UnmarshalPayload(body io.ReadCloser) (*VersionPayload, error) {
	payload := VersionPayload{}

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return nil, err
	}

	return &payload, nil
}

func (r *VersionResolver) UnmarshalResponse(body io.ReadCloser) (*OCPIVersionsResponse, error) {
	response := OCPIVersionsResponse{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
