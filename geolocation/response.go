package geolocation

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/util"
)

type GeoLocationPayload struct {
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Name      *string `json:"name,omitempty"`
}

func (r *GeoLocationPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewGeoLocationPayload(geoLocation db.GeoLocation) *GeoLocationPayload {
	return &GeoLocationPayload{
		Latitude:  geoLocation.Latitude,
		Longitude: geoLocation.Longitude,
		Name:      util.NilString(geoLocation.Name.String),
	}
}

func NewCreateGeoLocationParams(payload *GeoLocationPayload) db.CreateGeoLocationParams {
	return db.CreateGeoLocationParams{
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		Name:      util.SqlNullString(payload.Name),
	}
}

func NewUpdateGeoLocationParams(id int64, payload *GeoLocationPayload) db.UpdateGeoLocationParams {
	return db.UpdateGeoLocationParams{
		ID:        id,
		Latitude:  payload.Latitude,
		Longitude: payload.Longitude,
		Name:      util.SqlNullString(payload.Name),
	}
}

func (r *GeoLocationResolver) CreateGeoLocationPayload(ctx context.Context, geoLocation db.GeoLocation) *GeoLocationPayload {
	return NewGeoLocationPayload(geoLocation)
}

func (r *GeoLocationResolver) CreateGeoLocationListPayload(ctx context.Context, geoLocations []db.GeoLocation) []*GeoLocationPayload {
	list := []*GeoLocationPayload{}
	for _, geoLocation := range geoLocations {
		list = append(list, r.CreateGeoLocationPayload(ctx, geoLocation))
	}
	return list
}
