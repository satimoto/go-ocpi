package evse

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *EvseResolver) CreateCapabilityListDto(ctx context.Context, capabilities []db.Capability) []*string {
	list := []*string{}

	for i := 0; i < len(capabilities); i++ {
		list = append(list, &capabilities[i].Text)
	}

	return list
}

func (r *EvseResolver) CreateEvseDto(ctx context.Context, evse db.Evse) *dto.EvseDto {
	response := dto.NewEvseDto(evse)

	statusSchedules, err := r.Repository.ListStatusSchedules(ctx, evse.ID)

	if err != nil {
		metrics.RecordError("OCPI231", "Error listing status schedules", err)
		log.Printf("OCPI231: EvseID=%v", evse.ID)
	} else {
		response.StatusSchedule = r.CreateStatusScheduleListDto(ctx, statusSchedules)
	}

	capabilities, err := r.Repository.ListEvseCapabilities(ctx, evse.ID)

	if err != nil {
		metrics.RecordError("OCPI232", "Error listing evse capabilities", err)
		log.Printf("OCPI232: EvseID=%v", evse.ID)
	} else {
		response.Capabilities = r.CreateCapabilityListDto(ctx, capabilities)
	}

	connectors, err := r.Repository.ListConnectors(ctx, evse.ID)

	if err != nil {
		metrics.RecordError("OCPI233", "Error listing connectors", err)
		log.Printf("OCPI233: EvseID=%v", evse.ID)
	} else {
		response.Connectors = r.ConnectorResolver.CreateConnectorListDto(ctx, connectors)
	}

	if evse.GeoLocationID.Valid {
		geoLocation, err := r.Repository.GetGeoLocation(ctx, evse.GeoLocationID.Int64)

		if err != nil {
			metrics.RecordError("OCPI234", "Error listing connectors", err)
			log.Printf("OCPI234: GeoLocationID=%#v", evse.GeoLocationID)
		} else {
			response.Coordinates = r.GeoLocationResolver.CreateGeoLocationDto(ctx, geoLocation)
		}
	}

	directions, err := r.Repository.ListEvseDirections(ctx, evse.ID)

	if err != nil {
		metrics.RecordError("OCPI235", "Error listing evse directions", err)
		log.Printf("OCPI235: EvseID=%v", evse.ID)
	} else {
		response.Directions = r.DisplayTextResolver.CreateDisplayTextListDto(ctx, directions)
	}

	parkingRestrictions, err := r.Repository.ListEvseParkingRestrictions(ctx, evse.ID)

	if err != nil {
		metrics.RecordError("OCPI236", "Error listing evse parking restrictions", err)
		log.Printf("OCPI236: EvseID=%v", evse.ID)
	} else {
		response.ParkingRestrictions = r.CreateParkingRestrictionListDto(ctx, parkingRestrictions)
	}

	images, err := r.Repository.ListEvseImages(ctx, evse.ID)

	if err != nil {
		metrics.RecordError("OCPI237", "Error listing evse images", err)
		log.Printf("OCPI237: EvseID=%v", evse.ID)
	} else {
		response.Images = r.ImageResolver.CreateImageListDto(ctx, images)
	}

	return response
}

func (r *EvseResolver) CreateEvseListDto(ctx context.Context, evses []db.Evse) []*dto.EvseDto {
	list := []*dto.EvseDto{}

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

func (r *EvseResolver) CreateStatusScheduleDto(ctx context.Context, statusSchedule db.StatusSchedule) *coreDto.StatusScheduleDto {
	return coreDto.NewStatusScheduleDto(statusSchedule)
}

func (r *EvseResolver) CreateStatusScheduleListDto(ctx context.Context, statusSchedules []db.StatusSchedule) []*coreDto.StatusScheduleDto {
	list := []*coreDto.StatusScheduleDto{}

	for _, statusSchedule := range statusSchedules {
		list = append(list, r.CreateStatusScheduleDto(ctx, statusSchedule))
	}

	return list
}
