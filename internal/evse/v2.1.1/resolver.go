package evse

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/internal/image"
	"github.com/satimoto/go-ocpi-api/internal/util"
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

func (r *EvseResolver) ReplaceEvse(ctx context.Context, locationID int64, uid string, dto *EvseDto) *db.Evse {
	if dto != nil {
		evse, err := r.Repository.GetEvseByUid(ctx, uid)

		if err == nil {
			evseParams := NewUpdateEvseByUidParams(evse)

			if dto.Coordinates != nil {
				geoLocationID := util.SqlNullInt64(nil)
				geomPoint := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, dto.Coordinates)

				if geomPoint != nil {
					evseParams.Geom = *geomPoint
					evseParams.GeoLocationID = geoLocationID
				}
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
				geomPoint := r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, dto.Coordinates)

				if geomPoint != nil {
					evseParams.Geom = *geomPoint
					evseParams.GeoLocationID = geoLocationID
				}
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
