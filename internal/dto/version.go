package dto

import (
	"net/http"

	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type OcpiVersionsDto struct {
	Data          []*VersionDto `json:"data,omitempty"`
	StatusCode    int16         `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	Timestamp     ocpitype.Time `json:"timestamp"`
}

type VersionDto struct {
	Version *string `json:"version"`
	Url     *string `json:"url"`
}

func (r *VersionDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
