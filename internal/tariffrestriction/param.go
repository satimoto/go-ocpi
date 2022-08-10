package tariffrestriction

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateTariffRestrictionParams(dto *TariffRestrictionDto) db.CreateTariffRestrictionParams {
	return db.CreateTariffRestrictionParams{
		StartTime:  *dto.StartTime,
		EndTime:    *dto.EndTime,
		StartTime2: util.SqlNullString(dto.StartTime2),
		EndTime2:   util.SqlNullString(dto.EndTime2),
	}
}

func NewUpdateTariffRestrictionParams(id int64, dto *TariffRestrictionDto) db.UpdateTariffRestrictionParams {
	return db.UpdateTariffRestrictionParams{
		ID:         id,
		StartTime:  *dto.StartTime,
		EndTime:    *dto.EndTime,
		StartTime2: util.SqlNullString(dto.StartTime2),
		EndTime2:   util.SqlNullString(dto.EndTime2),
	}
}
