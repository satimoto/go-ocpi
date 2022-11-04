package evse

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/image"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *EvseResolver) ReplaceEvse(ctx context.Context, locationID int64, uid string, evseDto *dto.EvseDto) *db.Evse {
	if evseDto != nil {
		evse, err := r.Repository.GetEvseByUid(ctx, uid)

		if err == nil {
			evseParams := param.NewUpdateEvseByUidParams(evse)

			if evseDto.Coordinates != nil {
				geoLocationID := evse.GeoLocationID
				geometry := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, evseDto.Coordinates)

				if geometry != nil {
					evseParams.Geom = dbUtil.SqlNullGeometry4326(geometry)
					evseParams.GeoLocationID = geoLocationID
				}
			}

			if evseDto.Capabilities != nil {
				evseParams.IsRemoteCapable = dbUtil.StringsContainString(evseDto.Capabilities, "REMOTE_START_STOP_CAPABLE")
				evseParams.IsRfidCapable = dbUtil.StringsContainString(evseDto.Capabilities, "RFID_READER")
			}

			if evseDto.EvseID != nil {
				evseParams.EvseID = dbUtil.SqlNullString(evseDto.EvseID)
				evseParams.Identifier = dbUtil.SqlNullString(GetEvseIdentifier(evseDto))
			}

			if evseDto.FloorLevel != nil {
				evseParams.FloorLevel = dbUtil.SqlNullString(evseDto.FloorLevel)
			}

			if evseDto.LastUpdated != nil {
				evseParams.LastUpdated = evseDto.LastUpdated.Time()
			}

			if evseDto.PhysicalReference != nil {
				evseParams.PhysicalReference = dbUtil.SqlNullString(evseDto.PhysicalReference)
			}

			if evseDto.Status != nil {
				evseParams.Status = *evseDto.Status
			}

			if evseDto.Status != nil && evseDto.LastUpdated != nil && evseDto.Status != &evse.Status {
				evseStatusPeriodParams := param.NewCreateEvseStatusPeriodParams(evse, evseDto.LastUpdated.Time())
				_, err := r.Repository.CreateEvseStatusPeriod(ctx, evseStatusPeriodParams)

				if err != nil {
					metrics.RecordError("OCPI275", "Error creating evse status period", err)
					log.Printf("OCPI275: Params=%#v", evseStatusPeriodParams)
				}

				metricEvsesStatus.WithLabelValues(string(evse.Status)).Dec()
				metricEvsesStatus.WithLabelValues(string(*evseDto.Status)).Inc()
			}

			updatedEvse, err := r.Repository.UpdateEvseByUid(ctx, evseParams)

			if err != nil {
				metrics.RecordError("OCPI099", "Error updating evse", err)
				log.Printf("OCPI099: Params=%#v", evseParams)
				return nil
			}

			evse = updatedEvse
		} else {
			evseParams := NewCreateEvseParams(locationID, evseDto)

			if evseDto.Coordinates != nil {
				geoLocationID := dbUtil.SqlNullInt64(nil)
				geometry := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, evseDto.Coordinates)

				if geometry != nil {
					evseParams.Geom = dbUtil.SqlNullGeometry4326(geometry)
					evseParams.GeoLocationID = geoLocationID
				}
			}

			if evseDto.Capabilities != nil {
				evseParams.IsRemoteCapable = dbUtil.StringsContainString(evseDto.Capabilities, "REMOTE_START_STOP_CAPABLE")
				evseParams.IsRfidCapable = dbUtil.StringsContainString(evseDto.Capabilities, "RFID_READER")
			}

			evse, err = r.Repository.CreateEvse(ctx, evseParams)

			if err != nil {
				metrics.RecordError("OCPI100", "Error creating evse", err)
				log.Printf("OCPI100: Params=%#v", evseParams)
				return nil
			}

			// Metrics: Increment total evses
			metricEvsesTotal.Inc()
			metricEvsesStatus.WithLabelValues(string(evse.Status)).Inc()
		}

		if evseDto.Capabilities != nil {
			r.replaceCapabilities(ctx, evse.ID, evseDto)
		}

		if evseDto.Connectors != nil {
			r.replaceConnectors(ctx, evse, evseDto)
		}

		if evseDto.Directions != nil {
			r.replaceDirections(ctx, evse.ID, evseDto)
		}

		if evseDto.Images != nil {
			r.replaceImages(ctx, evse.ID, evseDto)
		}

		if evseDto.ParkingRestrictions != nil {
			r.replaceParkingRestrictions(ctx, evse.ID, evseDto)
		}

		if evseDto.StatusSchedule != nil {
			r.replaceStatusSchedule(ctx, evse.ID, evseDto)
		}

		return &evse
	}

	return nil
}

func (r *EvseResolver) ReplaceEvses(ctx context.Context, locationID int64, evseDto []*dto.EvseDto) {
	if evseDto != nil {
		if evses, err := r.Repository.ListEvses(ctx, locationID); err == nil {
			evseMap := make(map[string]db.Evse)

			for _, evse := range evses {
				evseMap[evse.Uid] = evse
			}

			for _, evse := range evseDto {
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
					metrics.RecordError("OCPI101", "Error updating evse", err)
					log.Printf("OCPI101: Params=%#v", evseParams)
				}

				metricEvsesStatus.WithLabelValues(string(evse.Status)).Dec()
				metricEvsesStatus.WithLabelValues(string(db.EvseStatusREMOVED)).Inc()
			}
		}

		r.updateLocationAvailability(ctx, locationID)
	}
}

func (r *EvseResolver) replaceCapabilities(ctx context.Context, evseID int64, evseDto *dto.EvseDto) {
	r.Repository.UnsetEvseCapabilities(ctx, evseID)

	if capabilities, err := r.Repository.ListCapabilities(ctx); err == nil {
		filteredCapabilities := []db.Capability{}

		for _, capability := range capabilities {
			if dbUtil.StringsContainString(evseDto.Capabilities, capability.Text) {
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
				metrics.RecordError("OCPI102", "Error setting evse capability", err)
				log.Printf("OCPI102: Params=%#v", setEvseCapabilityParams)
			}
		}
	}
}

func (r *EvseResolver) replaceConnectors(ctx context.Context, evse db.Evse, evseDto *dto.EvseDto) {
	r.ConnectorResolver.ReplaceConnectors(ctx, evse, evseDto.Connectors)
}

func (r *EvseResolver) replaceDirections(ctx context.Context, evseID int64, evseDto *dto.EvseDto) {
	r.Repository.DeleteEvseDirections(ctx, evseID)

	for _, directionDto := range evseDto.Directions {
		displayTextParams := displaytext.NewCreateDisplayTextParams(directionDto)
		displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams)

		if err != nil {
			metrics.RecordError("OCPI103", "Error creating display text", err)
			log.Printf("OCPI103: Params=%#v", displayTextParams)
			continue
		}

		setEvseDirectionParams := db.SetEvseDirectionParams{
			EvseID:        evseID,
			DisplayTextID: displayText.ID,
		}
		err = r.Repository.SetEvseDirection(ctx, setEvseDirectionParams)

		if err != nil {
			metrics.RecordError("OCPI104", "Error setting evse direction", err)
			log.Printf("OCPI104: Params=%#v", setEvseDirectionParams)
		}
	}
}

func (r *EvseResolver) replaceImages(ctx context.Context, evseID int64, evseDto *dto.EvseDto) {
	r.Repository.DeleteEvseImages(ctx, evseID)

	for _, imageDto := range evseDto.Images {
		imageParams := image.NewCreateImageParams(imageDto)

		if image, err := r.ImageResolver.Repository.CreateImage(ctx, imageParams); err == nil {
			setEvseImageParams := db.SetEvseImageParams{
				EvseID:  evseID,
				ImageID: image.ID,
			}
			err := r.Repository.SetEvseImage(ctx, setEvseImageParams)

			if err != nil {
				metrics.RecordError("OCPI106", "Error setting evse image", err)
				log.Printf("OCPI106: Params=%#v", setEvseImageParams)
			}
		}
	}
}

func (r *EvseResolver) replaceParkingRestrictions(ctx context.Context, evseID int64, evseDto *dto.EvseDto) {
	r.Repository.UnsetEvseParkingRestrictions(ctx, evseID)

	if parkingRestrictions, err := r.Repository.ListParkingRestrictions(ctx); err == nil {
		filteredParkingRestrictions := []db.ParkingRestriction{}

		for _, parkingRestriction := range parkingRestrictions {
			if dbUtil.StringsContainString(evseDto.ParkingRestrictions, parkingRestriction.Text) {
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
				metrics.RecordError("OCPI107", "Error setting evse parking restriction", err)
				log.Printf("OCPI107: Params=%#v", setEvseParkingRestrictionParams)
			}
		}
	}
}

func (r *EvseResolver) replaceStatusSchedule(ctx context.Context, evseID int64, evseDto *dto.EvseDto) {
	r.Repository.DeleteStatusSchedules(ctx, evseID)

	for _, statusScheduleDto := range evseDto.StatusSchedule {
		statusScheduleParams := NewCreateStatusScheduleParams(evseID, statusScheduleDto)
		_, err := r.Repository.CreateStatusSchedule(ctx, statusScheduleParams)

		if err != nil {
			metrics.RecordError("OCPI108", "Error creating status schedule", err)
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
			metrics.RecordError("OCPI109", "Error updating location availability", err)
			log.Printf("OCPI109: Params=%#v", updateLocationAvailabilityParams)
		}
	}
}
