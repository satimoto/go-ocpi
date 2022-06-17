package businessdetail

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *BusinessDetailResolver) ReplaceBusinessDetail(ctx context.Context, id *sql.NullInt64, dto *BusinessDetailDto) {
	if dto != nil {
		logoID := r.ImageResolver.CreateImage(ctx, dto.Logo)

		if id.Valid {
			businessDetailParams := NewUpdateBusinessDetailParams(id.Int64, dto)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)

			_, err := r.Repository.UpdateBusinessDetail(ctx, businessDetailParams)

			if err != nil {
				util.LogOnError("OCPI272", "Error updating business detail", err)
				log.Printf("OCPI272: Params=%#v", businessDetailParams)
			}	
		} else {
			businessDetailParams := NewCreateBusinessDetailParams(dto)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)

			businessDetail, err := r.Repository.CreateBusinessDetail(ctx, businessDetailParams)

			if err != nil {
				util.LogOnError("OCPI016", "Error creating business detail", err)
				log.Printf("OCPI016: Params=%#v", businessDetailParams)
				return
			}

			id.Scan(businessDetail.ID)
		}
	}
}
