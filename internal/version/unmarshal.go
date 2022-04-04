package version

import (
	"encoding/json"
	"io"
	"time"
)

type OCPIVersionsResponse struct {
	Data          []*VersionDto `json:"data,omitempty"`
	StatusCode    int16         `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	Timestamp     time.Time     `json:"timestamp"`
}

func (r *VersionResolver) UnmarshalDto(body io.ReadCloser) (*VersionDto, error) {
	dto := VersionDto{}

	if err := json.NewDecoder(body).Decode(&dto); err != nil {
		return nil, err
	}

	return &dto, nil
}

func (r *VersionResolver) UnmarshalResponse(body io.ReadCloser) (*OCPIVersionsResponse, error) {
	response := OCPIVersionsResponse{}

	if err := json.NewDecoder(body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
