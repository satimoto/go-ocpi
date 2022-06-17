package geolocation

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
)

type GeoLocationDto struct {
	Latitude  DecimalString  `json:"latitude"`
	Longitude DecimalString  `json:"longitude"`
}

func (r *GeoLocationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewGeoLocationDto(geoLocation db.GeoLocation) *GeoLocationDto {
	return &GeoLocationDto{
		Latitude:  NewDecimalString(geoLocation.Latitude),
		Longitude: NewDecimalString(geoLocation.Longitude),
	}
}

func (r *GeoLocationResolver) CreateGeoLocationDto(ctx context.Context, geoLocation db.GeoLocation) *GeoLocationDto {
	return NewGeoLocationDto(geoLocation)
}

func (r *GeoLocationResolver) CreateGeoLocationListDto(ctx context.Context, geoLocations []db.GeoLocation) []*GeoLocationDto {
	list := []*GeoLocationDto{}
	
	for _, geoLocation := range geoLocations {
		list = append(list, r.CreateGeoLocationDto(ctx, geoLocation))
	}

	return list
}
