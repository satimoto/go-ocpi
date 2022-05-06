package versiondetail

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *VersionDetailResolver) ReplaceVersionEndpoints(ctx context.Context, versionID int64, dto *VersionDetailDto) []*db.VersionEndpoint {
	endpoints := []*db.VersionEndpoint{}

	if dto != nil {
		r.Repository.DeleteVersionEndpoints(ctx, versionID)

		for _, endpointDto := range dto.Endpoints {
			endpointParams := NewCreateVersionEndpointParams(versionID, endpointDto)

			if endpoint, err := r.Repository.CreateVersionEndpoint(ctx, endpointParams); err == nil {
				endpoints = append(endpoints, &endpoint)
			}
		}

	}

	return endpoints
}

func (r *VersionDetailResolver) PullVersionEndpoints(ctx context.Context, url string, header transportation.OcpiRequestHeader, versionID int64) []*db.VersionEndpoint {
	if response, err := r.OcpiRequester.Do(http.MethodGet, url, header, nil); err == nil {
		ocpiResponse, err := r.UnmarshalPullDto(response.Body)
		response.Body.Close()

		if err == nil && ocpiResponse.StatusCode == transportation.STATUS_CODE_OK {
			return r.ReplaceVersionEndpoints(ctx, versionID, ocpiResponse.Data)
		}
	}

	return []*db.VersionEndpoint{}
}
