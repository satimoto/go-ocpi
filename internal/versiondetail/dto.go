package versiondetail

import (
	"context"
	"fmt"
	"os"

	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/version"
)

func (r *VersionDetailResolver) CreateEndpointDto(ctx context.Context, apiDomain string, identifier string) *coreDto.EndpointDto {
	return &coreDto.EndpointDto{
		Identifier: identifier,
		Url:        fmt.Sprintf("%s/%s/%s", apiDomain, version.VERSION_2_1_1, identifier),
	}
}

func (r *VersionDetailResolver) CreateVersionDetailDto(ctx context.Context) *coreDto.VersionDetailDto {
	apiDomain := os.Getenv("API_DOMAIN")

	var endpoints []*coreDto.EndpointDto
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "cdrs"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "credentials"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "commands"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "locations"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "sessions"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "tariffs"))
	endpoints = append(endpoints, r.CreateEndpointDto(ctx, apiDomain, "tokens"))

	return &coreDto.VersionDetailDto{
		Version:   version.VERSION_2_1_1,
		Endpoints: endpoints,
	}
}
