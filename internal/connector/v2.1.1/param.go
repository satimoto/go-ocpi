package connector

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/util"
)

func NewCreateConnectorParams(evse db.Evse, connectorDto *dto.ConnectorDto) db.CreateConnectorParams {
	return db.CreateConnectorParams{
		EvseID:             evse.ID,
		Uid:                *connectorDto.Id,
		Identifier:         dbUtil.SqlNullString(GetConnectorIdentifier(evse, connectorDto)),
		Standard:           *connectorDto.Standard,
		Format:             *connectorDto.Format,
		PowerType:          *connectorDto.PowerType,
		Voltage:            *connectorDto.Voltage,
		Amperage:           *connectorDto.Amperage,
		Wattage:            util.CalculateWattage(*connectorDto.PowerType, *connectorDto.Voltage, *connectorDto.Amperage),
		TariffID:           dbUtil.SqlNullString(connectorDto.TariffID),
		TermsAndConditions: dbUtil.SqlNullString(connectorDto.TermsAndConditions),
		LastUpdated:        *connectorDto.LastUpdated,
	}
}
