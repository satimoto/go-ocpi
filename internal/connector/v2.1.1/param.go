package connector

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateConnectorParams(evseID int64, dto *ConnectorDto) db.CreateConnectorParams {
	return db.CreateConnectorParams{
		EvseID:             evseID,
		Uid:                *dto.Id,
		Standard:           *dto.Standard,
		Format:             *dto.Format,
		PowerType:          *dto.PowerType,
		Voltage:            *dto.Voltage,
		Amperage:           *dto.Amperage,
		TariffID:           util.SqlNullString(dto.TariffID),
		TermsAndConditions: util.SqlNullString(dto.TermsAndConditions),
		LastUpdated:        *dto.LastUpdated,
	}
}