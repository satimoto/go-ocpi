package evse

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/image"
)

func (r *EvseResolver) ReplaceEvse(ctx context.Context, locationID int64, uid string, dto *EvseDto) *db.Evse {
	if dto != nil {
		evse, err := r.Repository.GetEvseByUid(ctx, uid)

		if err == nil {
			evseParams := NewUpdateEvseByUidParams(evse)

			if dto.Coordinates != nil {
				geoLocationID := util.SqlNullInt64(nil)
				geometry := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, dto.Coordinates)

				if geometry != nil {
					evseParams.Geom = util.SqlNullGeometry4326(geometry)
					evseParams.GeoLocationID = geoLocationID
				}
			}

			if dto.Capabilities != nil {
				evseParams.IsRemoteCapable = util.StringsContainString(dto.Capabilities, "REMOTE_START_STOP_CAPABLE")
				evseParams.IsRfidCapable = util.StringsContainString(dto.Capabilities, "RFID_READER")
			}

			if dto.EvseID != nil {
				evseParams.EvseID = util.SqlNullString(dto.EvseID)
			}

			if dto.FloorLevel != nil {
				evseParams.FloorLevel = util.SqlNullString(dto.FloorLevel)
			}

			if dto.LastUpdated != nil {
				evseParams.LastUpdated = *dto.LastUpdated
			}

			if dto.PhysicalReference != nil {
				evseParams.PhysicalReference = util.SqlNullString(dto.PhysicalReference)
			}

			if dto.Status != nil {
				evseParams.Status = *dto.Status
			}

			evse, err = r.Repository.UpdateEvseByUid(ctx, evseParams)
		} else {
			evseParams := NewCreateEvseParams(locationID, dto)

			if dto.Coordinates != nil {
				geoLocationID := util.SqlNullInt64(evse.GeoLocationID)
				geometry := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, dto.Coordinates)

				if geometry != nil {
					evseParams.Geom = util.SqlNullGeometry4326(geometry)
					evseParams.GeoLocationID = geoLocationID
				}
			}

			if dto.Capabilities != nil {
				evseParams.IsRemoteCapable = util.StringsContainString(dto.Capabilities, "REMOTE_START_STOP_CAPABLE")
				evseParams.IsRfidCapable = util.StringsContainString(dto.Capabilities, "RFID_READER")
			}

			evse, err = r.Repository.CreateEvse(ctx, evseParams)
		}

		if dto.Capabilities != nil {
			r.replaceCapabilities(ctx, evse.ID, dto)
		}

		if dto.Connectors != nil {
			r.replaceConnectors(ctx, evse.ID, dto)
		}

		if dto.Directions != nil {
			r.replaceDirections(ctx, evse.ID, dto)
		}

		if dto.Images != nil {
			r.replaceImages(ctx, evse.ID, dto)
		}

		if dto.ParkingRestrictions != nil {
			r.replaceParkingRestrictions(ctx, evse.ID, dto)
		}

		if dto.StatusSchedule != nil {
			r.replaceStatusSchedule(ctx, evse.ID, dto)
		}

		return &evse
	}

	return nil
}

func (r *EvseResolver) ReplaceEvses(ctx context.Context, locationID int64, dto []*EvseDto) {
	if dto != nil {
		if evses, err := r.Repository.ListEvses(ctx, locationID); err == nil {
			evseMap := make(map[string]db.Evse)

			for _, evse := range evses {
				evseMap[evse.Uid] = evse
			}

			for _, evse := range dto {
				if evse.Uid != nil {
					r.ReplaceEvse(ctx, locationID, *evse.Uid, evse)
					delete(evseMap, *evse.Uid)
				}
			}

			for _, evse := range evseMap {
				evseParams := db.NewUpdateEvseByUidParams(evse)
				evseParams.Status = db.EvseStatusREMOVED

				r.Repository.UpdateEvseByUid(ctx, evseParams)
			}
		}

		r.updateLocationAvailability(ctx, locationID)
	}
}

func (r *EvseResolver) replaceCapabilities(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.UnsetEvseCapabilities(ctx, evseID)

	if capabilities, err := r.Repository.ListCapabilities(ctx); err == nil {
		filteredCapabilities := []*db.Capability{}

		for _, capability := range capabilities {
			if util.StringsContainString(dto.Capabilities, capability.Text) {
				filteredCapabilities = append(filteredCapabilities, &capability)
			}
		}

		for _, capability := range filteredCapabilities {
			r.Repository.SetEvseCapability(ctx, db.SetEvseCapabilityParams{
				EvseID:       evseID,
				CapabilityID: capability.ID,
			})
		}
	}
}

func (r *EvseResolver) replaceConnectors(ctx context.Context, evseID int64, dto *EvseDto) {
	r.ConnectorResolver.ReplaceConnectors(ctx, evseID, dto.Connectors)
}

func (r *EvseResolver) replaceDirections(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.DeleteEvseDirections(ctx, evseID)

	for _, directionDto := range dto.Directions {
		displayTextParams := displaytext.NewCreateDisplayTextParams(directionDto)

		if displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams); err == nil {
			r.Repository.SetEvseDirection(ctx, db.SetEvseDirectionParams{
				EvseID:        evseID,
				DisplayTextID: displayText.ID,
			})
		}
	}
}

func (r *EvseResolver) replaceImages(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.DeleteEvseImages(ctx, evseID)

	for _, imageDto := range dto.Images {
		imageParams := image.NewCreateImageParams(imageDto)

		if image, err := r.ImageResolver.Repository.CreateImage(ctx, imageParams); err == nil {
			r.Repository.SetEvseImage(ctx, db.SetEvseImageParams{
				EvseID:  evseID,
				ImageID: image.ID,
			})
		}
	}
}

func (r *EvseResolver) replaceParkingRestrictions(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.UnsetEvseParkingRestrictions(ctx, evseID)

	if parkingRestrictions, err := r.Repository.ListParkingRestrictions(ctx); err == nil {
		filteredParkingRestrictions := []*db.ParkingRestriction{}

		for _, parkingRestriction := range parkingRestrictions {
			if util.StringsContainString(dto.ParkingRestrictions, parkingRestriction.Text) {
				filteredParkingRestrictions = append(filteredParkingRestrictions, &parkingRestriction)
			}
		}

		for _, parkingRestriction := range filteredParkingRestrictions {
			r.Repository.SetEvseParkingRestriction(ctx, db.SetEvseParkingRestrictionParams{
				EvseID:               evseID,
				ParkingRestrictionID: parkingRestriction.ID,
			})
		}
	}
}

func (r *EvseResolver) replaceStatusSchedule(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.DeleteStatusSchedules(ctx, evseID)

	for _, statusScheduleDto := range dto.StatusSchedule {
		statusScheduleParams := NewCreateStatusScheduleParams(evseID, statusScheduleDto)

		r.Repository.CreateStatusSchedule(ctx, statusScheduleParams)
	}
}

func (r *EvseResolver) updateLocationAvailability(ctx context.Context, locationID int64) {
	updateLocationAvailabilityParams := NewUpdateLocationAvailabilityParams(locationID)

	if evses, err := r.Repository.ListEvses(ctx, locationID); err == nil {
		for _, evse := range evses {
			updateLocationAvailabilityParams.TotalEvses++

			if evse.Status == db.EvseStatusAVAILABLE {
				updateLocationAvailabilityParams.AvailableEvses++
			}

			if evse.IsRemoteCapable {
				updateLocationAvailabilityParams.IsRemoteCapable = true
			}
			if evse.IsRfidCapable {
				updateLocationAvailabilityParams.IsRfidCapable = true
			}
		}

		r.Repository.UpdateLocationAvailability(ctx, updateLocationAvailabilityParams)
	}
}
