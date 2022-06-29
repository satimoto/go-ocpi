package location

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	"github.com/satimoto/go-ocpi/internal/energymix"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/geolocation"
	"github.com/satimoto/go-ocpi/internal/image"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
	"github.com/satimoto/go-ocpi/internal/openingtime"
)

type OcpiLocationsDto struct {
	Data          []*LocationDto `json:"data,omitempty"`
	StatusCode    int16          `json:"status_code"`
	StatusMessage string         `json:"status_message"`
	Timestamp     ocpitype.Time  `json:"timestamp"`
}

type LocationDto struct {
	ID                 *string                           `json:"id"`
	Type               *db.LocationType                  `json:"type"`
	Name               *string                           `json:"name,omitempty"`
	Address            *string                           `json:"address"`
	City               *string                           `json:"city"`
	PostalCode         *string                           `json:"postal_code"`
	Country            *string                           `json:"country"`
	Coordinates        *geolocation.GeoLocationDto       `json:"coordinates"`
	RelatedLocations   []*AdditionalGeoLocationDto       `json:"related_locations"`
	Evses              []*evse.EvseDto                   `json:"evses"`
	Directions         []*displaytext.DisplayTextDto     `json:"directions"`
	Facilities         []*string                         `json:"facilities"`
	Operator           *businessdetail.BusinessDetailDto `json:"operator,omitempty"`
	Suboperator        *businessdetail.BusinessDetailDto `json:"suboperator,omitempty"`
	Owner              *businessdetail.BusinessDetailDto `json:"owner,omitempty"`
	TimeZone           *string                           `json:"time_zone,omitempty"`
	OpeningTimes       *openingtime.OpeningTimeDto       `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool                             `json:"charging_when_closed"`
	Images             []*image.ImageDto                 `json:"images"`
	EnergyMix          *energymix.EnergyMixDto           `json:"energy_mix"`
	LastUpdated        *time.Time                        `json:"last_updated"`
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

func (r *LocationResolver) CreateFacilityListDto(ctx context.Context, facilities []db.Facility) []*string {
	list := []*string{}

	for i := 0; i < len(facilities); i++ {
		list = append(list, &facilities[i].Text)
	}

	return list
}

func (r *LocationResolver) CreateLocationDto(ctx context.Context, location db.Location) *LocationDto {
	response := NewLocationDto(location)

	geoLocation, err := r.GeoLocationResolver.Repository.GetGeoLocation(ctx, location.GeoLocationID)

	if err != nil {
		util.LogOnError("OCPI238", "Error retrieving geolocation", err)
		log.Printf("OCPI238: GeoLocationID=%v", location.GeoLocationID)
	} else {
		response.Coordinates = r.GeoLocationResolver.CreateGeoLocationDto(ctx, geoLocation)
	}

	additionalGeoLocations, err := r.Repository.ListAdditionalGeoLocations(ctx, location.ID)

	if err != nil {
		util.LogOnError("OCPI239", "Error listing additional geo locations", err)
		log.Printf("OCPI239: LocationID=%v", location.ID)
	} else {
		response.RelatedLocations = r.CreateAdditionalGeoLocationListDto(ctx, additionalGeoLocations)
	}

	evses, err := r.Repository.ListEvses(ctx, location.ID)

	if err != nil {
		util.LogOnError("OCPI240", "Error listing evses", err)
		log.Printf("OCPI240: LocationID=%v", location.ID)
	} else {
		response.Evses = r.EvseResolver.CreateEvseListDto(ctx, evses)
	}

	directions, err := r.Repository.ListLocationDirections(ctx, location.ID)

	if err != nil {
		util.LogOnError("OCPI241", "Error listing location directions", err)
		log.Printf("OCPI241: LocationID=%v", location.ID)
	} else {
		response.Directions = r.DisplayTextResolver.CreateDisplayTextListDto(ctx, directions)
	}

	facilities, err := r.Repository.ListLocationFacilities(ctx, location.ID)

	if err != nil {
		util.LogOnError("OCPI242", "Error listing location facilities", err)
		log.Printf("OCPI242: LocationID=%v", location.ID)
	} else {
		response.Facilities = r.CreateFacilityListDto(ctx, facilities)
	}

	if location.EnergyMixID.Valid {
		energyMix, err := r.EnergyMixResolver.Repository.GetEnergyMix(ctx, location.EnergyMixID.Int64)

		if err != nil {
			util.LogOnError("OCPI243", "Error retrieving energy mix", err)
			log.Printf("OCPI243: EnergyMixID=%#v", location.EnergyMixID)
		} else {
			response.EnergyMix = r.EnergyMixResolver.CreateEnergyMixDto(ctx, energyMix)
		}
	}

	if location.OperatorID.Valid {
		operator, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.OperatorID.Int64)

		if err != nil {
			util.LogOnError("OCPI244", "Error retrieving operator business detail", err)
			log.Printf("OCPI244: OperatorID=%#v", location.OperatorID)
		} else {
			response.Operator = r.BusinessDetailResolver.CreateBusinessDetailDto(ctx, operator)
		}
	}

	if location.SuboperatorID.Valid {
		suboperator, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.SuboperatorID.Int64)

		if err != nil {
			util.LogOnError("OCPI245", "Error retrieving suboperator business detail", err)
			log.Printf("OCPI245: SuboperatorID=%#v", location.SuboperatorID)
		} else {
			response.Suboperator = r.BusinessDetailResolver.CreateBusinessDetailDto(ctx, suboperator)
		}
	}

	if location.OwnerID.Valid {
		owner, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, location.OwnerID.Int64)

		if err != nil {
			util.LogOnError("OCPI246", "Error retrieving owner business detail", err)
			log.Printf("OCPI246: OwnerID=%#v", location.OwnerID)
		} else {
			response.Owner = r.BusinessDetailResolver.CreateBusinessDetailDto(ctx, owner)
		}
	}

	if location.OpeningTimeID.Valid {
		openingTime, err := r.OpeningTimeResolver.Repository.GetOpeningTime(ctx, location.OpeningTimeID.Int64)

		if err != nil {
			util.LogOnError("OCPI247", "Error retrieving opening time", err)
			log.Printf("OCPI247: OpeningTimeID=%#v", location.OpeningTimeID)
		} else {
			response.OpeningTimes = r.OpeningTimeResolver.CreateOpeningTimeDto(ctx, openingTime)
		}
	}

	images, err := r.Repository.ListLocationImages(ctx, location.ID)

	if err != nil {
		util.LogOnError("OCPI248", "Error listing location images", err)
		log.Printf("OCPI248: LocationID=%v", location.ID)
	} else {
		response.Images = r.ImageResolver.CreateImageListDto(ctx, images)
	}

	return response
}

func (r *LocationResolver) CreateLocationListDto(ctx context.Context, locations []db.Location) []render.Renderer {
	list := []render.Renderer{}

	for _, location := range locations {
		list = append(list, r.CreateLocationDto(ctx, location))
	}

	return list
}

type AdditionalGeoLocationDto struct {
	Latitude  ocpitype.String             `json:"latitude"`
	Longitude ocpitype.String             `json:"longitude"`
	Name      *displaytext.DisplayTextDto `json:"name,omitempty"`
}

func (r *AdditionalGeoLocationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewAdditionalGeoLocationDto(additionalGeoLocation db.AdditionalGeoLocation) *AdditionalGeoLocationDto {
	return &AdditionalGeoLocationDto{
		Latitude:  ocpitype.NewString(additionalGeoLocation.Latitude),
		Longitude: ocpitype.NewString(additionalGeoLocation.Longitude),
	}
}

func (r *LocationResolver) CreateAdditionalGeoLocationDto(ctx context.Context, additionalGeoLocation db.AdditionalGeoLocation) *AdditionalGeoLocationDto {
	response := NewAdditionalGeoLocationDto(additionalGeoLocation)

	if additionalGeoLocation.DisplayTextID.Valid {
		displayText, err := r.DisplayTextResolver.Repository.GetDisplayText(ctx, additionalGeoLocation.DisplayTextID.Int64)

		if err != nil {
			util.LogOnError("OCPI271", "Error retrieving display text", err)
			log.Printf("OCPI271: LocationID=%v, DisplayTextID=%v", additionalGeoLocation.LocationID, additionalGeoLocation.DisplayTextID.Int64)
		} else {
			response.Name = r.DisplayTextResolver.CreateDisplayTextDto(ctx, displayText)
		}
	}

	return response
}

func (r *LocationResolver) CreateAdditionalGeoLocationListDto(ctx context.Context, additionalGeoLocations []db.AdditionalGeoLocation) []*AdditionalGeoLocationDto {
	list := []*AdditionalGeoLocationDto{}

	for _, additionalGeoLocation := range additionalGeoLocations {
		list = append(list, r.CreateAdditionalGeoLocationDto(ctx, additionalGeoLocation))
	}

	return list
}
