package businessdetail

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/metric"
)

func (r *BusinessDetailResolver) CreateBusinessDetailDto(ctx context.Context, businessDetail db.BusinessDetail) *coreDto.BusinessDetailDto {
	response := coreDto.NewBusinessDetailDto(businessDetail)

	if businessDetail.LogoID.Valid {
		image, err := r.ImageResolver.Repository.GetImage(ctx, businessDetail.LogoID.Int64)

		if err != nil {
			metrics.RecordError("OCPI221", "Error retrieving image", err)
			log.Printf("OCPI222: LogoID=%#v", businessDetail.LogoID)
			return response
		}

		response.Logo = r.ImageResolver.CreateImageDto(ctx, image)
	}

	return response
}
