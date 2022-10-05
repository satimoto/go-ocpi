package businessdetail

import (
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateBusinessDetailParams(businessDetailDto *coreDto.BusinessDetailDto) db.CreateBusinessDetailParams {
	return db.CreateBusinessDetailParams{
		Name:    businessDetailDto.Name,
		Website: util.SqlNullString(businessDetailDto.Website),
	}
}

func NewUpdateBusinessDetailParams(id int64, businessDetailDto *coreDto.BusinessDetailDto) db.UpdateBusinessDetailParams {
	return db.UpdateBusinessDetailParams{
		ID:      id,
		Name:    businessDetailDto.Name,
		Website: util.SqlNullString(businessDetailDto.Website),
	}
}
