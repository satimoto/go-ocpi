package evse

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	"github.com/satimoto/go-ocpi/internal/image"
)

func (r *EvseResolver) ReplaceEvse(ctx context.Context, locationID int64, uid string, dto *EvseDto) *db.Evse {
	if dto != nil {
		evse, err := r.Repository.GetEvseByUid(ctx, uid)

		if err == nil {
			evseParams := param.NewUpdateEvseByUidParams(evse)

			if dto.Coordinates != nil {
				geoLocationID := dbUtil.SqlNullInt64(nil)
				geometry := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, dto.Coordinates)

				if geometry != nil {
					evseParams.Geom = dbUtil.SqlNullGeometry4326(geometry)
					evseParams.GeoLocationID = geoLocationID
				}
			}

			if dto.Capabilities != nil {
				evseParams.IsRemoteCapable = dbUtil.StringsContainString(dto.Capabilities, "REMOTE_START_STOP_CAPABLE")
				evseParams.IsRfidCapable = dbUtil.StringsContainString(dto.Capabilities, "RFID_READER")
			}

			if dto.EvseID != nil {
				evseParams.EvseID = dbUtil.SqlNullString(dto.EvseID)
				evseParams.Identifier = dbUtil.SqlNullString(GetEvseIdentifier(dto))
			}

			if dto.FloorLevel != nil {
				evseParams.FloorLevel = dbUtil.SqlNullString(dto.FloorLevel)
			}

			if dto.LastUpdated != nil {
				evseParams.LastUpdated = *dto.LastUpdated
			}

			if dto.PhysicalReference != nil {
				evseParams.PhysicalReference = dbUtil.SqlNullString(dto.PhysicalReference)
			}

			if dto.Status != nil {
				evseParams.Status = *dto.Status
			}

			updatedEvse, err := r.Repository.UpdateEvseByUid(ctx, evseParams)

			if err != nil {
				dbUtil.LogOnError("OCPI099", "Error updating evse", err)
				log.Printf("OCPI099: Params=%#v", evseParams)
				return nil
			}

			evse = updatedEvse
		} else {
			evseParams := NewCreateEvseParams(locationID, dto)

			if dto.Coordinates != nil {
				geoLocationID := dbUtil.SqlNullInt64(evse.GeoLocationID)
				geometry := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, dto.Coordinates)

				if geometry != nil {
					evseParams.Geom = dbUtil.SqlNullGeometry4326(geometry)
					evseParams.GeoLocationID = geoLocationID
				}
			}

			if dto.Capabilities != nil {
				evseParams.IsRemoteCapable = dbUtil.StringsContainString(dto.Capabilities, "REMOTE_START_STOP_CAPABLE")
				evseParams.IsRfidCapable = dbUtil.StringsContainString(dto.Capabilities, "RFID_READER")
			}

			evse, err = r.Repository.CreateEvse(ctx, evseParams)

			if err != nil {
				dbUtil.LogOnError("OCPI100", "Error creating evse", err)
				log.Printf("OCPI100: Params=%#v", evseParams)
				return nil
			}
		}

		if dto.Capabilities != nil {
			r.replaceCapabilities(ctx, evse.ID, dto)
		}

		if dto.Connectors != nil {
			r.replaceConnectors(ctx, evse, dto)
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
				evseParams := param.NewUpdateEvseByUidParams(evse)
				evseParams.Status = db.EvseStatusREMOVED
				_, err := r.Repository.UpdateEvseByUid(ctx, evseParams)

				if err != nil {
					dbUtil.LogOnError("OCPI101", "Error updating evse", err)
					log.Printf("OCPI101: Params=%#v", evseParams)
				}
			}
		}

		r.updateLocationAvailability(ctx, locationID)
	}
}

func (r *EvseResolver) replaceCapabilities(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.UnsetEvseCapabilities(ctx, evseID)

	if capabilities, err := r.Repository.ListCapabilities(ctx); err == nil {
		filteredCapabilities := []db.Capability{}

		for _, capability := range capabilities {
			if dbUtil.StringsContainString(dto.Capabilities, capability.Text) {
				filteredCapabilities = append(filteredCapabilities, capability)
			}
		}

		for _, capability := range filteredCapabilities {
			setEvseCapabilityParams := db.SetEvseCapabilityParams{
				EvseID:       evseID,
				CapabilityID: capability.ID,
			}
			err := r.Repository.SetEvseCapability(ctx, setEvseCapabilityParams)

			if err != nil {
				dbUtil.LogOnError("OCPI102", "Error setting evse capability", err)
				log.Printf("OCPI102: Params=%#v", setEvseCapabilityParams)
			}
		}
	}
}

func (r *EvseResolver) replaceConnectors(ctx context.Context, evse db.Evse, dto *EvseDto) {
	r.ConnectorResolver.ReplaceConnectors(ctx, evse, dto.Connectors)
}

func (r *EvseResolver) replaceDirections(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.DeleteEvseDirections(ctx, evseID)

	for _, directionDto := range dto.Directions {
		displayTextParams := displaytext.NewCreateDisplayTextParams(directionDto)
		displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams)

		if err != nil {
			dbUtil.LogOnError("OCPI103", "Error creating display text", err)
			log.Printf("OCPI103: Params=%#v", displayTextParams)
			continue
		}

		setEvseDirectionParams := db.SetEvseDirectionParams{
			EvseID:        evseID,
			DisplayTextID: displayText.ID,
		}
		err = r.Repository.SetEvseDirection(ctx, setEvseDirectionParams)

		if err != nil {
			dbUtil.LogOnError("OCPI104", "Error setting evse direction", err)
			log.Printf("OCPI104: Params=%#v", setEvseDirectionParams)
		}
	}
}

func (r *EvseResolver) replaceImages(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.DeleteEvseImages(ctx, evseID)

	for _, imageDto := range dto.Images {
		imageParams := image.NewCreateImageParams(imageDto)

		if image, err := r.ImageResolver.Repository.CreateImage(ctx, imageParams); err == nil {
			setEvseImageParams := db.SetEvseImageParams{
				EvseID:  evseID,
				ImageID: image.ID,
			}
			err := r.Repository.SetEvseImage(ctx, setEvseImageParams)

			if err != nil {
				dbUtil.LogOnError("OCPI106", "Error setting evse image", err)
				log.Printf("OCPI106: Params=%#v", setEvseImageParams)
			}
		}
	}
}

func (r *EvseResolver) replaceParkingRestrictions(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.UnsetEvseParkingRestrictions(ctx, evseID)

	if parkingRestrictions, err := r.Repository.ListParkingRestrictions(ctx); err == nil {
		filteredParkingRestrictions := []db.ParkingRestriction{}

		for _, parkingRestriction := range parkingRestrictions {
			if dbUtil.StringsContainString(dto.ParkingRestrictions, parkingRestriction.Text) {
				filteredParkingRestrictions = append(filteredParkingRestrictions, parkingRestriction)
			}
		}

		for _, parkingRestriction := range filteredParkingRestrictions {
			setEvseParkingRestrictionParams := db.SetEvseParkingRestrictionParams{
				EvseID:               evseID,
				ParkingRestrictionID: parkingRestriction.ID,
			}
			err := r.Repository.SetEvseParkingRestriction(ctx, setEvseParkingRestrictionParams)

			if err != nil {
				dbUtil.LogOnError("OCPI107", "Error setting evse parking restriction", err)
				log.Printf("OCPI107: Params=%#v", setEvseParkingRestrictionParams)
			}
		}
	}
}

func (r *EvseResolver) replaceStatusSchedule(ctx context.Context, evseID int64, dto *EvseDto) {
	r.Repository.DeleteStatusSchedules(ctx, evseID)

	for _, statusScheduleDto := range dto.StatusSchedule {
		statusScheduleParams := NewCreateStatusScheduleParams(evseID, statusScheduleDto)
		_, err := r.Repository.CreateStatusSchedule(ctx, statusScheduleParams)

		if err != nil {
			dbUtil.LogOnError("OCPI108", "Error creating status schedule", err)
			log.Printf("OCPI108: Params=%#v", statusScheduleParams)
		}
	}
}

func (r *EvseResolver) updateLocationAvailability(ctx context.Context, locationID int64) {
	updateLocationAvailabilityParams := param.NewUpdateLocationAvailabilityParams(locationID)

	if evses, err := r.Repository.ListEvses(ctx, locationID); err == nil {
		for _, evse := range evses {
			if evse.Status != db.EvseStatusREMOVED {
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
		}

		err := r.Repository.UpdateLocationAvailability(ctx, updateLocationAvailabilityParams)

		if err != nil {
			dbUtil.LogOnError("OCPI109", "Error updating location availability", err)
			log.Printf("OCPI109: Params=%#v", updateLocationAvailabilityParams)
		}
	}
}
