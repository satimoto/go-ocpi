package connector

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dtUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/util"
)

func NewCreateConnectorParams(evseID int64, dto *ConnectorDto) db.CreateConnectorParams {
	wattage := util.CalculateWattage(*dto.PowerType, *dto.Voltage, *dto.Amperage)

	return db.CreateConnectorParams{
		EvseID:             evseID,
		Uid:                *dto.Id,
		Standard:           *dto.Standard,
		Format:             *dto.Format,
		PowerType:          *dto.PowerType,
		Voltage:            *dto.Voltage,
		Amperage:           *dto.Amperage,
		Wattage:            wattage,
		TariffID:           dtUtil.SqlNullString(dto.TariffID),
		TermsAndConditions: dtUtil.SqlNullString(dto.TermsAndConditions),
		LastUpdated:        *dto.LastUpdated,
	}
}
