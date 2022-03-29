package evse

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/internal/image"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func (r *EvseResolver) CreateCapabilityListPayload(ctx context.Context, capabilities []db.Capability) []*string {
	list := []*string{}
	for i := 0; i < len(capabilities); i++ {
		list = append(list, &capabilities[i].Text)
	}
	return list
}

type EvsePayload struct {
	Uid                 *string                           `json:"uid"`
	EvseID              *string                           `json:"evse_id,omitempty"`
	Status              *db.EvseStatus                    `json:"status"`
	StatusSchedule      []*StatusSchedulePayload          `json:"status_schedule"`
	Capabilities        []*string                         `json:"capabilities"`
	Connectors          []*connector.ConnectorPayload     `json:"connectors"`
	FloorLevel          *string                           `json:"floor_level,omitempty"`
	Coordinates         *geolocation.GeoLocationPayload   `json:"coordinates,omitempty"`
	PhysicalReference   *string                           `json:"physical_reference,omitempty"`
	Directions          []*displaytext.DisplayTextPayload `json:"directions"`
	ParkingRestrictions []*string                         `json:"parking_restrictions"`
	Images              []*image.ImagePayload             `json:"images"`
	LastUpdated         *time.Time                        `json:"last_updated"`
}

func (r *EvsePayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEvsePayload(evse db.Evse) *EvsePayload {
	return &EvsePayload{
		Uid:               &evse.Uid,
		EvseID:            util.NilString(evse.EvseID.String),
		Status:            &evse.Status,
		FloorLevel:        util.NilString(evse.FloorLevel.String),
		PhysicalReference: util.NilString(evse.PhysicalReference.String),
		LastUpdated:       &evse.LastUpdated,
	}
}

func NewCreateEvseParams(locationID int64, payload *EvsePayload) db.CreateEvseParams {
	return db.CreateEvseParams{
		Uid:               *payload.Uid,
		EvseID:            util.SqlNullString(payload.EvseID),
		LocationID:        locationID,
		Status:            *payload.Status,
		FloorLevel:        util.SqlNullString(payload.FloorLevel),
		PhysicalReference: util.SqlNullString(payload.PhysicalReference),
		LastUpdated:       *payload.LastUpdated,
	}
}

func NewUpdateEvseByUidParams(evse db.Evse) db.UpdateEvseByUidParams {
	return db.UpdateEvseByUidParams{
		Uid:               evse.Uid,
		EvseID:            evse.EvseID,
		Status:            evse.Status,
		FloorLevel:        evse.FloorLevel,
		Geom:              evse.Geom,
		GeoLocationID:     evse.GeoLocationID,
		PhysicalReference: evse.PhysicalReference,
		LastUpdated:       evse.LastUpdated,
	}
}

func (r *EvseResolver) CreateEvsePayload(ctx context.Context, evse db.Evse) *EvsePayload {
	response := NewEvsePayload(evse)

	if statusSchedules, err := r.Repository.ListStatusSchedules(ctx, evse.ID); err == nil {
		response.StatusSchedule = r.CreateStatusScheduleListPayload(ctx, statusSchedules)
	}

	if capabilities, err := r.Repository.ListEvseCapabilities(ctx, evse.ID); err == nil {
		response.Capabilities = r.CreateCapabilityListPayload(ctx, capabilities)
	}

	if connectors, err := r.Repository.ListConnectors(ctx, evse.ID); err == nil {
		response.Connectors = r.ConnectorResolver.CreateConnectorListPayload(ctx, connectors)
	}

	if evse.GeoLocationID.Valid {
		if geoLocation, err := r.Repository.GetGeoLocation(ctx, evse.GeoLocationID.Int64); err == nil {
			response.Coordinates = r.GeoLocationResolver.CreateGeoLocationPayload(ctx, geoLocation)
		}
	}

	if directions, err := r.Repository.ListEvseDirections(ctx, evse.ID); err == nil {
		response.Directions = r.DisplayTextResolver.CreateDisplayTextListPayload(ctx, directions)
	}

	if parkingRestrictions, err := r.Repository.ListEvseParkingRestrictions(ctx, evse.ID); err == nil {
		response.ParkingRestrictions = r.CreateParkingRestrictionListPayload(ctx, parkingRestrictions)
	}

	if images, err := r.Repository.ListEvseImages(ctx, evse.ID); err == nil {
		response.Images = r.ImageResolver.CreateImageListPayload(ctx, images)
	}

	return response
}

func (r *EvseResolver) CreateEvseListPayload(ctx context.Context, evses []db.Evse) []*EvsePayload {
	list := []*EvsePayload{}
	for _, evse := range evses {
		list = append(list, r.CreateEvsePayload(ctx, evse))
	}
	return list
}

func (r *EvseResolver) CreateParkingRestrictionListPayload(ctx context.Context, parkingRestrictions []db.ParkingRestriction) []*string {
	list := []*string{}
	for _, parkingRestriction := range parkingRestrictions {
		list = append(list, &parkingRestriction.Text)
	}
	return list
}

type StatusSchedulePayload struct {
	PeriodBegin *time.Time    `json:"period_begin"`
	PeriodEnd   *time.Time    `json:"period_end,omitempty"`
	Status      db.EvseStatus `json:"status"`
}

func (r *StatusSchedulePayload) Render(writer http.ResponseWriter, request *http.Request) error {
	if r.PeriodEnd.IsZero() {
		r.PeriodEnd = nil
	}
	return nil
}

func NewStatusSchedulePayload(statusSchedule db.StatusSchedule) *StatusSchedulePayload {
	return &StatusSchedulePayload{
		PeriodBegin: &statusSchedule.PeriodBegin,
		PeriodEnd:   util.NilTime(statusSchedule.PeriodEnd.Time),
		Status:      statusSchedule.Status,
	}
}

func NewCreateStatusScheduleParams(evseID int64, payload *StatusSchedulePayload) db.CreateStatusScheduleParams {
	return db.CreateStatusScheduleParams{
		EvseID:      evseID,
		PeriodBegin: *payload.PeriodBegin,
		PeriodEnd:   util.SqlNullTime(payload.PeriodEnd),
	}
}

func (r *EvseResolver) CreateStatusSchedulePayload(ctx context.Context, statusSchedule db.StatusSchedule) *StatusSchedulePayload {
	return NewStatusSchedulePayload(statusSchedule)
}

func (r *EvseResolver) CreateStatusScheduleListPayload(ctx context.Context, statusSchedules []db.StatusSchedule) []*StatusSchedulePayload {
	list := []*StatusSchedulePayload{}
	for _, statusSchedule := range statusSchedules {
		list = append(list, r.CreateStatusSchedulePayload(ctx, statusSchedule))
	}
	return list
}
