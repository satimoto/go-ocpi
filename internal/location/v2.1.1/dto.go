package location

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/internal/image"
	"github.com/satimoto/go-ocpi-api/internal/openingtime"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type OCPILocationsDto struct {
	Data          []*LocationDto `json:"data,omitempty"`
	StatusCode    int16          `json:"status_code"`
	StatusMessage string         `json:"status_message"`
	Timestamp     time.Time      `json:"timestamp"`
}

type LocationDto struct {
	ID                 *string                           `json:"id"`
	Type               *db.LocationType                  `json:"type"`
	Name               *string                           `json:"name,omitempty"`
	Address            *string                           `json:"address"`
	City               *string                           `json:"city"`
	PostalCode         *string                           `json:"postal_code"`
	Country            *string                           `json:"country"`
	Coordinates        *geolocation.GeoLocationDto       `json:"coordinates"`
	RelatedLocations   []*geolocation.GeoLocationDto     `json:"related_locations"`
	Evses              []*evse.EvseDto                   `json:"evses"`
	Directions         []*displaytext.DisplayTextDto     `json:"directions"`
	Facilities         []*string                         `json:"facilities"`
	Operator           *businessdetail.BusinessDetailDto `json:"operator,omitempty"`
	Suboperator        *businessdetail.BusinessDetailDto `json:"suboperator,omitempty"`
	Owner              *businessdetail.BusinessDetailDto `json:"owner,omitempty"`
	TimeZone           *string                           `json:"time_zone,omitempty"`
	OpeningTimes       *openingtime.OpeningTimeDto       `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                             `json:"charging_when_closed"`
	Images             []*image.ImageDto                 `json:"images"`
	EnergyMix          *energymix.EnergyMixDto           `json:"energy_mix"`
	LastUpdated        *time.Time                        `json:"last_updated"`
}

func (r *LocationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewLocationDto(location db.Location) *LocationDto {
	return &LocationDto{
		ID:                 &location.Uid,
		Type:               &location.Type,
		Name:               util.NilString(location.Name),
		Address:            &location.Address,
		City:               &location.City,
		PostalCode:         &location.PostalCode,
		Country:            &location.Country,
		TimeZone:           util.NilString(location.TimeZone),
		ChargingWhenClosed: &location.ChargingWhenClosed,
		LastUpdated:        &location.LastUpdated,
	}
}

func (r *LocationResolver) CreateFacilityListDto(ctx context.Context, facilities []db.Facility) []*string {
	list := []*string{}
	for i := 0; i < len(facilities); i++ {
		list = append(list, &facilities[i].Text)
	}
	return list
}

func (r *LocationResolver) CreateLocationDto(ctx context.Context, location db.Location) *LocationDto {
	response := NewLocationDto(location)

	if geoLocation, err := r.GeoLocationResolver.Repository.GetGeoLocation(ctx, location.GeoLocationID); err == nil {
		response.Coordinates = r.GeoLocationResolver.CreateGeoLocationDto(ctx, geoLocation)
	}

	if relatedLocations, err := r.Repository.ListRelatedLocations(ctx, location.ID); err == nil {
		response.RelatedLocations = r.GeoLocationResolver.CreateGeoLocationListDto(ctx, relatedLocations)
	}

	if evses, err := r.Repository.ListEvses(ctx, location.ID); err == nil {
		response.Evses = r.EvseResolver.CreateEvseListDto(ctx, evses)
	}

	if directions, err := r.Repository.ListLocationDirections(ctx, location.ID); err == nil {
		response.Directions = r.DisplayTextResolver.CreateDisplayTextListDto(ctx, directions)
	}

	if facilities, err := r.Repository.ListLocationFacilities(ctx, location.ID); err == nil {
		response.Facilities = r.CreateFacilityListDto(ctx, facilities)
	}

	if location.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixResolver.Repository.GetEnergyMix(ctx, location.EnergyMixID.Int64); err == nil {
			response.EnergyMix = r.EnergyMixResolver.CreateEnergyMixDto(ctx, energyMix)
		}
	}

	if location.OperatorID.Valid {
		if operator, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.OperatorID.Int64); err == nil {
			response.Operator = r.BusinessDetailResolver.CreateBusinessDetailDto(ctx, operator)
		}
	}

	if location.SuboperatorID.Valid {
		if suboperator, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.SuboperatorID.Int64); err == nil {
			response.Suboperator = r.BusinessDetailResolver.CreateBusinessDetailDto(ctx, suboperator)
		}
	}

	if location.OwnerID.Valid {
		if owner, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.OwnerID.Int64); err == nil {
			response.Owner = r.BusinessDetailResolver.CreateBusinessDetailDto(ctx, owner)
		}
	}

	if location.OpeningTimeID.Valid {
		if openingTime, err := r.OpeningTimeResolver.Repository.GetOpeningTime(ctx, location.OpeningTimeID.Int64); err == nil {
			response.OpeningTimes = r.OpeningTimeResolver.CreateOpeningTimeDto(ctx, openingTime)
		}
	}

	if images, err := r.Repository.ListLocationImages(ctx, location.ID); err == nil {
		response.Images = r.ImageResolver.CreateImageListDto(ctx, images)
	}

	return response
}

func (r *LocationResolver) CreateLocationListDto(ctx context.Context, locations []db.Location) []render.Renderer {
	list := []render.Renderer{}
	for _, location := range locations {
		list = append(list, r.CreateLocationDto(ctx, location))
	}
	return list
}
