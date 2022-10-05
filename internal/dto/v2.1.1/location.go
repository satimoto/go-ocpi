package dto

import (
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type OcpiLocationsDto struct {
	Data          []*LocationDto `json:"data,omitempty"`
	StatusCode    int16          `json:"status_code"`
	StatusMessage string         `json:"status_message"`
	Timestamp     ocpitype.Time  `json:"timestamp"`
}

type LocationDto struct {
	ID                 *string                             `json:"id"`
	Type               *db.LocationType                    `json:"type"`
	Name               *string                             `json:"name,omitempty"`
	Address            *string                             `json:"address"`
	City               *string                             `json:"city"`
	PostalCode         *string                             `json:"postal_code"`
	Country            *string                             `json:"country"`
	Coordinates        *coreDto.GeoLocationDto             `json:"coordinates"`
	RelatedLocations   []*coreDto.AdditionalGeoLocationDto `json:"related_locations"`
	Evses              []*EvseDto                          `json:"evses"`
	Directions         []*coreDto.DisplayTextDto           `json:"directions"`
	Facilities         []*string                           `json:"facilities"`
	Operator           *coreDto.BusinessDetailDto          `json:"operator,omitempty"`
	Suboperator        *coreDto.BusinessDetailDto          `json:"suboperator,omitempty"`
	Owner              *coreDto.BusinessDetailDto          `json:"owner,omitempty"`
	TimeZone           *string                             `json:"time_zone,omitempty"`
	OpeningTimes       *coreDto.OpeningTimeDto             `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                               `json:"charging_when_closed"`
	Images             []*coreDto.ImageDto                 `json:"images"`
	EnergyMix          *coreDto.EnergyMixDto               `json:"energy_mix"`
	LastUpdated        *time.Time                          `json:"last_updated"`
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
