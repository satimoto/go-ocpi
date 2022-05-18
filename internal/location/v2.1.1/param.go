package location

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
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
		AvailableEvses:     0,
		TotalEvses:         0,
		IsRemoteCapable:    false,
		IsRfidCapable:      false,
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
		AvailableEvses:     location.AvailableEvses,
		TotalEvses:         location.TotalEvses,
		IsRemoteCapable:    location.IsRemoteCapable,
		IsRfidCapable:      location.IsRfidCapable,
		OperatorID:         location.OperatorID,
		SuboperatorID:      location.SuboperatorID,
		OwnerID:            location.OwnerID,
		TimeZone:           location.TimeZone,
		ChargingWhenClosed: location.ChargingWhenClosed,
		EnergyMixID:        location.EnergyMixID,
		LastUpdated:        location.LastUpdated,
	}
}
