package tariff

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func NewCreateTariffParams(tariffDto *dto.TariffDto) db.CreateTariffParams {
	return db.CreateTariffParams{
		Uid:          *tariffDto.ID,
		Currency:     *tariffDto.Currency,
		TariffAltUrl: util.SqlNullString(tariffDto.TariffAltUrl),
		LastUpdated:  *tariffDto.LastUpdated,
	}
}
