package evse

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	connector "github.com/satimoto/go-ocpi-api/connector/v2.1.1"
	"github.com/satimoto/go-ocpi-api/displaytext"
	"github.com/satimoto/go-ocpi-api/geolocation"
	"github.com/satimoto/go-ocpi-api/image"
	"github.com/satimoto/go-ocpi-api/util"
)

type EvseRepository interface {
	CreateEvse(ctx context.Context, arg db.CreateEvseParams) (db.Evse, error)
	CreateStatusSchedule(ctx context.Context, arg db.CreateStatusScheduleParams) (db.StatusSchedule, error)
	DeleteConnectors(ctx context.Context, evseID int64) error
	DeleteEvseDirections(ctx context.Context, evseID int64) error
	DeleteEvseImages(ctx context.Context, evseID int64) error
	DeleteStatusSchedules(ctx context.Context, evseID int64) error
	GetEvse(ctx context.Context, id int64) (db.Evse, error)
	GetEvseByUid(ctx context.Context, uid string) (db.Evse, error)
	GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error)
	ListCapabilities(ctx context.Context) ([]db.Capability, error)
	ListConnectors(ctx context.Context, evseID int64) ([]db.Connector, error)
	ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error)
	ListEvseCapabilities(ctx context.Context, evseID int64) ([]db.Capability, error)
	ListEvseDirections(ctx context.Context, evseID int64) ([]db.DisplayText, error)
	ListEvseImages(ctx context.Context, evseID int64) ([]db.Image, error)
	ListEvseParkingRestrictions(ctx context.Context, evseID int64) ([]db.ParkingRestriction, error)
	ListParkingRestrictions(ctx context.Context) ([]db.ParkingRestriction, error)
	ListStatusSchedules(ctx context.Context, evseID int64) ([]db.StatusSchedule, error)
	SetEvseCapability(ctx context.Context, arg db.SetEvseCapabilityParams) error
	SetEvseDirection(ctx context.Context, arg db.SetEvseDirectionParams) error
	SetEvseImage(ctx context.Context, arg db.SetEvseImageParams) error
	SetEvseParkingRestriction(ctx context.Context, arg db.SetEvseParkingRestrictionParams) error
	UnsetEvseCapabilities(ctx context.Context, evseID int64) error
	UnsetEvseParkingRestrictions(ctx context.Context, evseID int64) error
	UpdateEvseByUid(ctx context.Context, arg db.UpdateEvseByUidParams) (db.Evse, error)
	UpdateEvseLastUpdated(ctx context.Context, arg db.UpdateEvseLastUpdatedParams) error
	UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error
}

type EvseResolver struct {
	Repository EvseRepository
	*connector.ConnectorResolver
	*displaytext.DisplayTextResolver
	*geolocation.GeoLocationResolver
	*image.ImageResolver
}

func NewResolver(repositoryService *db.RepositoryService) *EvseResolver {
	repo := EvseRepository(repositoryService)
	return &EvseResolver{
		Repository:          repo,
		ConnectorResolver:   connector.NewResolver(repositoryService),
		DisplayTextResolver: displaytext.NewResolver(repositoryService),
		GeoLocationResolver: geolocation.NewResolver(repositoryService),
		ImageResolver:       image.NewResolver(repositoryService),
	}
}

func (r *EvseResolver) ReplaceEvse(ctx context.Context, locationID int64, uid string, payload *EvsePayload) *db.Evse {
	if payload != nil {
		evse, err := r.Repository.GetEvseByUid(ctx, uid)

		if err == nil {
			evseParams := NewUpdateEvseByUidParams(evse)

			if payload.Coordinates != nil {
				geoLocationID := util.SqlNullInt64(nil)
				geomPoint := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, payload.Coordinates)

				if geomPoint != nil {
					evseParams.Geom = *geomPoint
					evseParams.GeoLocationID = geoLocationID
				}
			}

			if payload.EvseID != nil {
				evseParams.EvseID = util.SqlNullString(payload.EvseID)
			}

			if payload.FloorLevel != nil {
				evseParams.FloorLevel = util.SqlNullString(payload.FloorLevel)
			}

			if payload.LastUpdated != nil {
				evseParams.LastUpdated = *payload.LastUpdated
			}

			if payload.PhysicalReference != nil {
				evseParams.PhysicalReference = util.SqlNullString(payload.PhysicalReference)
			}

			if payload.Status != nil {
				evseParams.Status = *payload.Status
			}

			evse, err = r.Repository.UpdateEvseByUid(ctx, evseParams)
		} else {
			evseParams := NewCreateEvseParams(locationID, payload)

			if payload.Coordinates != nil {
				geoLocationID := util.SqlNullInt64(evse.GeoLocationID)
				geomPoint := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, payload.Coordinates)

				if geomPoint != nil {
					evseParams.Geom = *geomPoint
					evseParams.GeoLocationID = geoLocationID
				}
			}

			evse, err = r.Repository.CreateEvse(ctx, evseParams)
		}

		if payload.Capabilities != nil {
			r.replaceCapabilities(ctx, evse.ID, payload)
		}

		if payload.Connectors != nil {
			r.replaceConnectors(ctx, evse.ID, payload)
		}

		if payload.Directions != nil {
			r.replaceDirections(ctx, evse.ID, payload)
		}

		if payload.Images != nil {
			r.replaceImages(ctx, evse.ID, payload)
		}

		if payload.ParkingRestrictions != nil {
			r.replaceParkingRestrictions(ctx, evse.ID, payload)
		}

		if payload.StatusSchedule != nil {
			r.replaceStatusSchedule(ctx, evse.ID, payload)
		}

		return &evse
	}

	return nil
}

func (r *EvseResolver) ReplaceEvses(ctx context.Context, locationID int64, payload []*EvsePayload) {
	if payload != nil {
		if evses, err := r.Repository.ListEvses(ctx, locationID); err == nil {
			evseMap := make(map[string]db.Evse)

			for _, evse := range evses {
				evseMap[evse.Uid] = evse
			}

			for _, evse := range payload {
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
	}
}

func (r *EvseResolver) replaceCapabilities(ctx context.Context, evseID int64, payload *EvsePayload) {
	r.Repository.UnsetEvseCapabilities(ctx, evseID)

	if capabilities, err := r.Repository.ListCapabilities(ctx); err == nil {
		filteredCapabilities := []*db.Capability{}

		for _, capability := range capabilities {
			if util.StringsContainString(payload.Capabilities, capability.Text) {
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

func (r *EvseResolver) replaceConnectors(ctx context.Context, evseID int64, payload *EvsePayload) {
	r.ConnectorResolver.ReplaceConnectors(ctx, evseID, payload.Connectors)
}

func (r *EvseResolver) replaceDirections(ctx context.Context, evseID int64, payload *EvsePayload) {
	r.Repository.DeleteEvseDirections(ctx, evseID)

	for _, directionPayload := range payload.Directions {
		displayTextParams := displaytext.NewCreateDisplayTextParams(directionPayload)

		if displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams); err == nil {
			r.Repository.SetEvseDirection(ctx, db.SetEvseDirectionParams{
				EvseID:        evseID,
				DisplayTextID: displayText.ID,
			})
		}
	}
}

func (r *EvseResolver) replaceImages(ctx context.Context, evseID int64, payload *EvsePayload) {
	r.Repository.DeleteEvseImages(ctx, evseID)

	for _, imagePayload := range payload.Images {
		imageParams := image.NewCreateImageParams(imagePayload)

		if image, err := r.ImageResolver.Repository.CreateImage(ctx, imageParams); err == nil {
			r.Repository.SetEvseImage(ctx, db.SetEvseImageParams{
				EvseID:  evseID,
				ImageID: image.ID,
			})
		}
	}
}

func (r *EvseResolver) replaceParkingRestrictions(ctx context.Context, evseID int64, payload *EvsePayload) {
	r.Repository.UnsetEvseParkingRestrictions(ctx, evseID)

	if parkingRestrictions, err := r.Repository.ListParkingRestrictions(ctx); err == nil {
		filteredParkingRestrictions := []*db.ParkingRestriction{}

		for _, parkingRestriction := range parkingRestrictions {
			if util.StringsContainString(payload.ParkingRestrictions, parkingRestriction.Text) {
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

func (r *EvseResolver) replaceStatusSchedule(ctx context.Context, evseID int64, payload *EvsePayload) {
	r.Repository.DeleteStatusSchedules(ctx, evseID)

	for _, statusSchedulePayload := range payload.StatusSchedule {
		statusScheduleParams := NewCreateStatusScheduleParams(evseID, statusSchedulePayload)

		r.Repository.CreateStatusSchedule(ctx, statusScheduleParams)
	}
}
