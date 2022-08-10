package businessdetail

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
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
