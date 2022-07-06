package connector

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/util"
)

func NewCreateConnectorParams(evse db.Evse, dto *ConnectorDto) db.CreateConnectorParams {
	params := db.CreateConnectorParams{
		EvseID:             evse.ID,
		Uid:                *dto.Id,
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

	if evse.EvseID.Valid {
		params.ConnectorID = dbUtil.SqlNullString(evse.EvseID.String + *dto.Id)
	}

	return params
}
