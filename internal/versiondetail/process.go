package versiondetail

import (
	"context"
	"log"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *VersionDetailResolver) ReplaceVersionEndpoints(ctx context.Context, versionID int64, dto *VersionDetailDto) []*db.VersionEndpoint {
	endpoints := []*db.VersionEndpoint{}

	if dto != nil {
		r.Repository.DeleteVersionEndpoints(ctx, versionID)

		for _, endpointDto := range dto.Endpoints {
			endpointParams := NewCreateVersionEndpointParams(versionID, endpointDto)
			endpoint, err := r.Repository.CreateVersionEndpoint(ctx, endpointParams)

			if err != nil {
				util.LogOnError("OCPI216", "Error creating version endpoint", err)
				log.Printf("OCPI216: Params=%#v", endpointParams)
				continue
			}

			endpoints = append(endpoints, &endpoint)
		}
	}

	return endpoints
}

func (r *VersionDetailResolver) PullVersionEndpoints(ctx context.Context, url string, header transportation.OcpiRequestHeader, versionID int64) []*db.VersionEndpoint {
	response, err := r.OcpiRequester.Do(http.MethodGet, url, header, nil)

	if err != nil {
		util.LogOnError("OCPI217", "Error making request", err)
		log.Printf("OCPI217: Method=%v, Url=%v, Header=%#v", http.MethodGet, url, header)
		return []*db.VersionEndpoint{}
	}

	pullDto, err := r.UnmarshalPullDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		util.LogOnError("OCPI218", "Error unmarshalling response", err)
		util.LogHttpResponse("OCPI218", url, response, true)
		return []*db.VersionEndpoint{}
	}

	if pullDto.StatusCode != transportation.STATUS_CODE_OK {
		util.LogOnError("OCPI219", "Error response failure", err)
		util.LogHttpRequest("OCPI219", url, response.Request, true)
		util.LogHttpResponse("OCPI219", url, response, true)
		log.Printf("OCPI219: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
		return []*db.VersionEndpoint{}
	}

	return r.ReplaceVersionEndpoints(ctx, versionID, pullDto.Data)
}
