package businessdetail

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *BusinessDetailResolver) ReplaceBusinessDetail(ctx context.Context, id *sql.NullInt64, businessDetailDto *coreDto.BusinessDetailDto) {
	if businessDetailDto != nil {
		logoID := r.ImageResolver.CreateImage(ctx, businessDetailDto.Logo)

		if id.Valid {
			businessDetailParams := NewUpdateBusinessDetailParams(id.Int64, businessDetailDto)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)

			_, err := r.Repository.UpdateBusinessDetail(ctx, businessDetailParams)

			if err != nil {
				metrics.RecordError("OCPI272", "Error updating business detail", err)
				log.Printf("OCPI272: Params=%#v", businessDetailParams)
			}
		} else {
			businessDetailParams := NewCreateBusinessDetailParams(businessDetailDto)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)

			businessDetail, err := r.Repository.CreateBusinessDetail(ctx, businessDetailParams)

			if err != nil {
				metrics.RecordError("OCPI016", "Error creating business detail", err)
				log.Printf("OCPI016: Params=%#v", businessDetailParams)
				return
			}

			id.Scan(businessDetail.ID)
		}
	}
}
