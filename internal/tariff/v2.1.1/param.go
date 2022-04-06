package tariff

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateTariffParams(dto *TariffDto) db.CreateTariffParams {
	return db.CreateTariffParams{
		Uid:          *dto.ID,
		Currency:     *dto.Currency,
		TariffAltUrl: util.SqlNullString(dto.TariffAltUrl),
		LastUpdated:  *dto.LastUpdated,
	}
}

func NewUpdateTariffByUidParams(tariff db.Tariff) db.UpdateTariffByUidParams {
	return db.UpdateTariffByUidParams{
		Uid:                 tariff.Uid,
		CountryCode:         tariff.CountryCode,
		PartyID:             tariff.PartyID,
		Currency:            tariff.Currency,
		TariffAltUrl:        tariff.TariffAltUrl,
		EnergyMixID:         tariff.EnergyMixID,
		TariffRestrictionID: tariff.TariffRestrictionID,
		LastUpdated:         tariff.LastUpdated,
	}
}
