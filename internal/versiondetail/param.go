package versiondetail

import (
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateVersionEndpointParams(versionID int64, endpointDto *coreDto.EndpointDto) db.CreateVersionEndpointParams {
	return db.CreateVersionEndpointParams{
		VersionID:  versionID,
		Identifier: endpointDto.Identifier,
		Url:        endpointDto.Url,
	}
}
