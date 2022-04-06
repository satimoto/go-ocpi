package geolocation

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type GeoLocationDto struct {
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Name      *string `json:"name,omitempty"`
}

func (r *GeoLocationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewGeoLocationDto(geoLocation db.GeoLocation) *GeoLocationDto {
	return &GeoLocationDto{
		Latitude:  geoLocation.Latitude,
		Longitude: geoLocation.Longitude,
		Name:      util.NilString(geoLocation.Name),
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
