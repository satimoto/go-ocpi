package versiondetail

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/satimoto/go-datastore/db"
)

var (
	version = "2.1.1"
)

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

func NewCreateVersionEndpointParams(versionID int64, dto *EndpointDto) db.CreateVersionEndpointParams {
	return db.CreateVersionEndpointParams{
		VersionID:  versionID,
		Identifier: dto.Identifier,
		Url:        dto.Url,
	}
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
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "locations"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "sessions"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "tariffs"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "tokens"))

	return &VersionDetailDto{
		Version:   version,
		Endpoints: endpoints,
	}
}
