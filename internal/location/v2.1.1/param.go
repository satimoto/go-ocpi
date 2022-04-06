package location

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateLocationParams(dto *LocationDto) db.CreateLocationParams {
	return db.CreateLocationParams{
		Uid:                *dto.ID,
		Type:               *dto.Type,
		Name:               util.SqlNullString(dto.Name),
		Address:            *dto.Address,
		City:               *dto.City,
		PostalCode:         *dto.PostalCode,
		Country:            *dto.Country,
		TimeZone:           util.SqlNullString(dto.TimeZone),
		ChargingWhenClosed: *dto.ChargingWhenClosed,
		LastUpdated:        *dto.LastUpdated,
	}
}

func NewUpdateLocationByUidParams(location db.Location) db.UpdateLocationByUidParams {
	return db.UpdateLocationByUidParams{
		Uid:                location.Uid,
		CountryCode:        location.CountryCode,
		PartyID:            location.PartyID,
		Type:               location.Type,
		Name:               location.Name,
		Address:            location.Address,
		City:               location.City,
		PostalCode:         location.PostalCode,
		Country:            location.Country,
		Geom:               location.Geom,
		GeoLocationID:      location.GeoLocationID,
		OperatorID:         location.OperatorID,
		SuboperatorID:      location.SuboperatorID,
		OwnerID:            location.OwnerID,
		TimeZone:           location.TimeZone,
		ChargingWhenClosed: location.ChargingWhenClosed,
		EnergyMixID:        location.EnergyMixID,
		LastUpdated:        location.LastUpdated,
	}
}
