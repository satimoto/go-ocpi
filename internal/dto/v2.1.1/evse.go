package dto

import (
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

type EvseDto struct {
	Uid                 *string                      `json:"uid"`
	EvseID              *string                      `json:"evse_id,omitempty"`
	Status              *db.EvseStatus               `json:"status"`
	StatusSchedule      []*coreDto.StatusScheduleDto `json:"status_schedule"`
	Capabilities        []*string                    `json:"capabilities"`
	Connectors          []*ConnectorDto              `json:"connectors"`
	FloorLevel          *string                      `json:"floor_level,omitempty"`
	Coordinates         *coreDto.GeoLocationDto      `json:"coordinates,omitempty"`
	PhysicalReference   *string                      `json:"physical_reference,omitempty"`
	Directions          []*coreDto.DisplayTextDto    `json:"directions"`
	ParkingRestrictions []*string                    `json:"parking_restrictions"`
	Images              []*coreDto.ImageDto            `json:"images"`
	LastUpdated         *time.Time                   `json:"last_updated"`
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
