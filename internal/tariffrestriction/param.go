package tariffrestriction

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateTariffRestrictionParams(tariffRestrictionDto *coreDto.TariffRestrictionDto) db.CreateTariffRestrictionParams {
	return db.CreateTariffRestrictionParams{
		StartTime:  *tariffRestrictionDto.StartTime,
		EndTime:    *tariffRestrictionDto.EndTime,
		StartTime2: util.SqlNullString(tariffRestrictionDto.StartTime2),
		EndTime2:   util.SqlNullString(tariffRestrictionDto.EndTime2),
	}
}

func NewUpdateTariffRestrictionParams(id int64, tariffRestrictionDto *coreDto.TariffRestrictionDto) db.UpdateTariffRestrictionParams {
	return db.UpdateTariffRestrictionParams{
		ID:         id,
		StartTime:  *tariffRestrictionDto.StartTime,
		EndTime:    *tariffRestrictionDto.EndTime,
		StartTime2: util.SqlNullString(tariffRestrictionDto.StartTime2),
		EndTime2:   util.SqlNullString(tariffRestrictionDto.EndTime2),
	}
}
