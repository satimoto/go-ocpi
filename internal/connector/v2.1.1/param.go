package connector

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/util"
)

func NewCreateConnectorParams(evse db.Evse, dto *ConnectorDto) db.CreateConnectorParams {
	return db.CreateConnectorParams{
		EvseID:             evse.ID,
		Uid:                *dto.Id,
		Identifier:         dbUtil.SqlNullString(GetConnectorIdentifier(evse, dto)),
		Standard:           *dto.Standard,
		Format:             *dto.Format,
		PowerType:          *dto.PowerType,
		Voltage:            *dto.Voltage,
		Amperage:           *dto.Amperage,
		Wattage:            util.CalculateWattage(*dto.PowerType, *dto.Voltage, *dto.Amperage),
		TariffID:           dbUtil.SqlNullString(dto.TariffID),
		TermsAndConditions: dbUtil.SqlNullString(dto.TermsAndConditions),
		LastUpdated:        *dto.LastUpdated,
	}
}
