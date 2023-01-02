package dto

import (
	"net/http"

	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type OcpiVersionDetailDto struct {
	Data          *VersionDetailDto `json:"data,omitempty"`
	StatusCode    int16             `json:"status_code"`
	StatusMessage string            `json:"status_message"`
	Timestamp     ocpitype.Time     `json:"timestamp"`
}

type VersionDetailDto struct {
	Version   string         `json:"version"`
	Endpoints []*EndpointDto `json:"endpoints"`
}

func (r *VersionDetailDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type EndpointDto struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
}
