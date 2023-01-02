package geolocation

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *GeoLocationResolver) CreateGeoLocationDto(ctx context.Context, geoLocation db.GeoLocation) *coreDto.GeoLocationDto {
	return coreDto.NewGeoLocationDto(geoLocation)
}

func (r *GeoLocationResolver) CreateGeoLocationListDto(ctx context.Context, geoLocations []db.GeoLocation) []*coreDto.GeoLocationDto {
	list := []*coreDto.GeoLocationDto{}

	for _, geoLocation := range geoLocations {
		list = append(list, r.CreateGeoLocationDto(ctx, geoLocation))
	}

	return list
}
