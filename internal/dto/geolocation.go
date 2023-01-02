package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type GeoLocationDto struct {
	Latitude  ocpitype.String `json:"latitude"`
	Longitude ocpitype.String `json:"longitude"`
}

func (r *GeoLocationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewGeoLocationDto(geoLocation db.GeoLocation) *GeoLocationDto {
	return &GeoLocationDto{
		Latitude:  ocpitype.NewString(geoLocation.Latitude),
		Longitude: ocpitype.NewString(geoLocation.Longitude),
	}
}
