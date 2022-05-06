package versiondetail

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	version = "2.1.1"
)

type OcpiVersionDetailDto struct {
	Data          *VersionDetailDto `json:"data,omitempty"`
	StatusCode    int16             `json:"status_code"`
	StatusMessage string            `json:"status_message"`
	Timestamp     time.Time         `json:"timestamp"`
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

func (r *VersionDetailResolver) CreateEndpointDto(ctx context.Context, apiDomain string, identifier string) *EndpointDto {
	return &EndpointDto{
		Identifier: identifier,
		Url:        fmt.Sprintf("%s/%s/%s", apiDomain, version, identifier),
	}
}

func (r *VersionDetailResolver) CreateVersionDetailDto(ctx context.Context) *VersionDetailDto {
	apiDomain := os.Getenv("API_DOMAIN")

	var endpoints []*EndpointDto
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "cdrs"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "credentials"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "commands"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "locations"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "sessions"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "tariffs"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "tokens"))

	return &VersionDetailDto{
		Version:   version,
		Endpoints: endpoints,
	}
}
