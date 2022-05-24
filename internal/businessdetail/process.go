package businessdetail

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *BusinessDetailResolver) ReplaceBusinessDetail(ctx context.Context, id *int64, dto *BusinessDetailDto) {
	if dto != nil {
		logoID := r.ImageResolver.CreateImage(ctx, dto.Logo)

		if id == nil {
			businessDetailParams := NewCreateBusinessDetailParams(dto)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)
			
			businessDetail, err := r.Repository.CreateBusinessDetail(ctx, businessDetailParams)

			if err != nil {
				util.LogOnError("OCPI016", "Error creating business detail", err)
				log.Printf("OCPI016: Params=%#v", businessDetailParams)
				return
			}

			id = &businessDetail.ID
		} else {
			businessDetailParams := NewUpdateBusinessDetailParams(*id, dto)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)

			r.Repository.UpdateBusinessDetail(ctx, businessDetailParams)
		}
	}
}
