package versiondetail

import (
	"github.com/satimoto/go-datastore/pkg/db"
)

func NewCreateVersionEndpointParams(versionID int64, dto *EndpointDto) db.CreateVersionEndpointParams {
	return db.CreateVersionEndpointParams{
		VersionID:  versionID,
		Identifier: dto.Identifier,
		Url:        dto.Url,
	}
}
