package evse

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/internal/image"
)

func (r *EvseResolver) CreateCapabilityListDto(ctx context.Context, capabilities []db.Capability) []*string {
	list := []*string{}
	for i := 0; i < len(capabilities); i++ {
		list = append(list, &capabilities[i].Text)
	}
	return list
}

type EvseDto struct {
	Uid                 *string                       `json:"uid"`
	EvseID              *string                       `json:"evse_id,omitempty"`
	Status              *db.EvseStatus                `json:"status"`
	StatusSchedule      []*StatusScheduleDto          `json:"status_schedule"`
	Capabilities        []*string                     `json:"capabilities"`
	Connectors          []*connector.ConnectorDto     `json:"connectors"`
	FloorLevel          *string                       `json:"floor_level,omitempty"`
	Coordinates         *geolocation.GeoLocationDto   `json:"coordinates,omitempty"`
	PhysicalReference   *string                       `json:"physical_reference,omitempty"`
	Directions          []*displaytext.DisplayTextDto `json:"directions"`
	ParkingRestrictions []*string                     `json:"parking_restrictions"`
	Images              []*image.ImageDto             `json:"images"`
	LastUpdated         *time.Time                    `json:"last_updated"`
}

func (r *EvseDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEvseDto(evse db.Evse) *EvseDto {
	return &EvseDto{
		Uid:               &evse.Uid,
		EvseID:            util.NilString(evse.EvseID),
		Status:            &evse.Status,
		FloorLevel:        util.NilString(evse.FloorLevel),
		PhysicalReference: util.NilString(evse.PhysicalReference),
		LastUpdated:       &evse.LastUpdated,
	}
}

func (r *EvseResolver) CreateEvseDto(ctx context.Context, evse db.Evse) *EvseDto {
	response := NewEvseDto(evse)

	if statusSchedules, err := r.Repository.ListStatusSchedules(ctx, evse.ID); err == nil {
		response.StatusSchedule = r.CreateStatusScheduleListDto(ctx, statusSchedules)
	}

	if capabilities, err := r.Repository.ListEvseCapabilities(ctx, evse.ID); err == nil {
		response.Capabilities = r.CreateCapabilityListDto(ctx, capabilities)
	}

	if connectors, err := r.Repository.ListConnectors(ctx, evse.ID); err == nil {
		response.Connectors = r.ConnectorResolver.CreateConnectorListDto(ctx, connectors)
	}

	if evse.GeoLocationID.Valid {
		if geoLocation, err := r.Repository.GetGeoLocation(ctx, evse.GeoLocationID.Int64); err == nil {
			response.Coordinates = r.GeoLocationResolver.CreateGeoLocationDto(ctx, geoLocation)
		}
	}

	if directions, err := r.Repository.ListEvseDirections(ctx, evse.ID); err == nil {
		response.Directions = r.DisplayTextResolver.CreateDisplayTextListDto(ctx, directions)
	}

	if parkingRestrictions, err := r.Repository.ListEvseParkingRestrictions(ctx, evse.ID); err == nil {
		response.ParkingRestrictions = r.CreateParkingRestrictionListDto(ctx, parkingRestrictions)
	}

	if images, err := r.Repository.ListEvseImages(ctx, evse.ID); err == nil {
		response.Images = r.ImageResolver.CreateImageListDto(ctx, images)
	}

	return response
}

func (r *EvseResolver) CreateEvseListDto(ctx context.Context, evses []db.Evse) []*EvseDto {
	list := []*EvseDto{}
	for _, evse := range evses {
		list = append(list, r.CreateEvseDto(ctx, evse))
	}
	return list
}

func (r *EvseResolver) CreateParkingRestrictionListDto(ctx context.Context, parkingRestrictions []db.ParkingRestriction) []*string {
	list := []*string{}
	for _, parkingRestriction := range parkingRestrictions {
		text := parkingRestriction.Text
		list = append(list, &text)
	}
	return list
}

type StatusScheduleDto struct {
	PeriodBegin *time.Time    `json:"period_begin"`
	PeriodEnd   *time.Time    `json:"period_end,omitempty"`
	Status      db.EvseStatus `json:"status"`
}

func (r *StatusScheduleDto) Render(writer http.ResponseWriter, request *http.Request) error {
	if r.PeriodEnd.IsZero() {
		r.PeriodEnd = nil
	}
	return nil
}

func NewStatusScheduleDto(statusSchedule db.StatusSchedule) *StatusScheduleDto {
	return &StatusScheduleDto{
		PeriodBegin: &statusSchedule.PeriodBegin,
		PeriodEnd:   util.NilTime(statusSchedule.PeriodEnd.Time),
		Status:      statusSchedule.Status,
	}
}

func (r *EvseResolver) CreateStatusScheduleDto(ctx context.Context, statusSchedule db.StatusSchedule) *StatusScheduleDto {
	return NewStatusScheduleDto(statusSchedule)
}

func (r *EvseResolver) CreateStatusScheduleListDto(ctx context.Context, statusSchedules []db.StatusSchedule) []*StatusScheduleDto {
	list := []*StatusScheduleDto{}
	for _, statusSchedule := range statusSchedules {
		list = append(list, r.CreateStatusScheduleDto(ctx, statusSchedule))
	}
	return list
}
