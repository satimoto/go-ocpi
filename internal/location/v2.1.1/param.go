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
