package tariff

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateTariffParams(dto *TariffDto) db.CreateTariffParams {
	return db.CreateTariffParams{
		Uid:          *dto.ID,
		Currency:     *dto.Currency,
		TariffAltUrl: util.SqlNullString(dto.TariffAltUrl),
		LastUpdated:  *dto.LastUpdated,
	}
}
