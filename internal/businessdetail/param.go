package businessdetail

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateBusinessDetailParams(dto *BusinessDetailDto) db.CreateBusinessDetailParams {
	return db.CreateBusinessDetailParams{
		Name:    dto.Name,
		Website: util.SqlNullString(dto.Website),
	}
}

func NewUpdateBusinessDetailParams(id int64, dto *BusinessDetailDto) db.UpdateBusinessDetailParams {
	return db.UpdateBusinessDetailParams{
		ID:      id,
		Name:    dto.Name,
		Website: util.SqlNullString(dto.Website),
	}
}
