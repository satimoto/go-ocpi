package location

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/businessdetail"
	"github.com/satimoto/go-ocpi-api/displaytext"
	"github.com/satimoto/go-ocpi-api/energymix"
	evse "github.com/satimoto/go-ocpi-api/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/geolocation"
	"github.com/satimoto/go-ocpi-api/image"
	"github.com/satimoto/go-ocpi-api/openingtime"
	"github.com/satimoto/go-ocpi-api/util"
)

func (r *LocationResolver) CreateFacilityListPayload(ctx context.Context, facilities []db.Facility) []*string {
	list := []*string{}
	for i := 0; i < len(facilities); i++ {
		list = append(list, &facilities[i].Text)
	}
	return list
}

type LocationPayload struct {
	ID                 *string                               `json:"id"`
	Type               *db.LocationType                      `json:"type"`
	Name               *string                               `json:"name,omitempty"`
	Address            *string                               `json:"address"`
	City               *string                               `json:"city"`
	PostalCode         *string                               `json:"postal_code"`
	Country            *string                               `json:"country"`
	Coordinates        *geolocation.GeoLocationPayload       `json:"coordinates"`
	RelatedLocations   []*geolocation.GeoLocationPayload     `json:"related_locations"`
	Evses              []*evse.EvsePayload                   `json:"evses"`
	Directions         []*displaytext.DisplayTextPayload     `json:"directions"`
	Facilities         []*string                             `json:"facilities"`
	Operator           *businessdetail.BusinessDetailPayload `json:"operator,omitempty"`
	Suboperator        *businessdetail.BusinessDetailPayload `json:"suboperator,omitempty"`
	Owner              *businessdetail.BusinessDetailPayload `json:"owner,omitempty"`
	TimeZone           *string                               `json:"time_zone,omitempty"`
	OpeningTimes       *openingtime.OpeningTimePayload       `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                                 `json:"charging_when_closed"`
	Images             []*image.ImagePayload                 `json:"images"`
	EnergyMix          *energymix.EnergyMixPayload           `json:"energy_mix"`
	LastUpdated        *time.Time                            `json:"last_updated"`
}

func (r *LocationPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewLocationPayload(location db.Location) *LocationPayload {
	return &LocationPayload{
		ID:                 &location.Uid,
		Type:               &location.Type,
		Name:               util.NilString(location.Name.String),
		Address:            &location.Address,
		City:               &location.City,
		PostalCode:         &location.PostalCode,
		Country:            &location.Country,
		TimeZone:           util.NilString(location.TimeZone.String),
		ChargingWhenClosed: &location.ChargingWhenClosed,
		LastUpdated:        &location.LastUpdated,
	}
}

func NewCreateLocationParams(payload *LocationPayload) db.CreateLocationParams {
	return db.CreateLocationParams{
		Uid:                *payload.ID,
		Type:               *payload.Type,
		Name:               util.SqlNullString(payload.Name),
		Address:            *payload.Address,
		City:               *payload.City,
		PostalCode:         *payload.PostalCode,
		Country:            *payload.Country,
		TimeZone:           util.SqlNullString(payload.TimeZone),
		ChargingWhenClosed: *payload.ChargingWhenClosed,
		LastUpdated:        *payload.LastUpdated,
	}
}

func NewUpdateLocationByUidParams(location db.Location) db.UpdateLocationByUidParams {
	return db.UpdateLocationByUidParams{
		Uid:                location.Uid,
		Type:               location.Type,
		Name:               location.Name,
		Address:            location.Address,
		City:               location.City,
		PostalCode:         location.PostalCode,
		Country:            location.Country,
		Geom:               location.Geom,
		GeoLocationID:      location.GeoLocationID,
		OperatorID:         location.OperatorID,
		SuboperatorID:      location.SuboperatorID,
		OwnerID:            location.OwnerID,
		TimeZone:           location.TimeZone,
		ChargingWhenClosed: location.ChargingWhenClosed,
		EnergyMixID:        location.EnergyMixID,
		LastUpdated:        location.LastUpdated,
	}
}

func (r *LocationResolver) CreateLocationPayload(ctx context.Context, location db.Location) *LocationPayload {
	response := NewLocationPayload(location)

	if geoLocation, err := r.GeoLocationResolver.Repository.GetGeoLocation(ctx, location.GeoLocationID); err == nil {
		response.Coordinates = r.GeoLocationResolver.CreateGeoLocationPayload(ctx, geoLocation)
	}

	if relatedLocations, err := r.Repository.ListRelatedLocations(ctx, location.ID); err == nil {
		response.RelatedLocations = r.GeoLocationResolver.CreateGeoLocationListPayload(ctx, relatedLocations)
	}

	if evses, err := r.Repository.ListEvses(ctx, location.ID); err == nil {
		response.Evses = r.EvseResolver.CreateEvseListPayload(ctx, evses)
	}

	if directions, err := r.Repository.ListLocationDirections(ctx, location.ID); err == nil {
		response.Directions = r.DisplayTextResolver.CreateDisplayTextListPayload(ctx, directions)
	}

	if facilities, err := r.Repository.ListLocationFacilities(ctx, location.ID); err == nil {
		response.Facilities = r.CreateFacilityListPayload(ctx, facilities)
	}

	if location.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixResolver.Repository.GetEnergyMix(ctx, location.EnergyMixID.Int64); err == nil {
			response.EnergyMix = r.EnergyMixResolver.CreateEnergyMixPayload(ctx, energyMix)
		}
	}

	if location.OperatorID.Valid {
		if operator, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.OperatorID.Int64); err == nil {
			response.Operator = r.BusinessDetailResolver.CreateBusinessDetailPayload(ctx, operator)
		}
	}

	if location.SuboperatorID.Valid {
		if suboperator, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.SuboperatorID.Int64); err == nil {
			response.Suboperator = r.BusinessDetailResolver.CreateBusinessDetailPayload(ctx, suboperator)
		}
	}

	if location.OwnerID.Valid {
		if owner, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.OwnerID.Int64); err == nil {
			response.Owner = r.BusinessDetailResolver.CreateBusinessDetailPayload(ctx, owner)
		}
	}

	if location.OpeningTimeID.Valid {
		if openingTime, err := r.OpeningTimeResolver.Repository.GetOpeningTime(ctx, location.OpeningTimeID.Int64); err == nil {
			response.OpeningTimes = r.OpeningTimeResolver.CreateOpeningTimePayload(ctx, openingTime)
		}
	}

	if images, err := r.Repository.ListLocationImages(ctx, location.ID); err == nil {
		response.Images = r.ImageResolver.CreateImageListPayload(ctx, images)
	}

	return response
}

func (r *LocationResolver) CreateLocationListPayload(ctx context.Context, locations []db.Location) []render.Renderer {
	list := []render.Renderer{}
	for _, location := range locations {
		list = append(list, r.CreateLocationPayload(ctx, location))
	}
	return list
}
