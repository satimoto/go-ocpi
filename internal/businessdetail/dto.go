package businessdetail

import (
	"context"
	"log"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/image"
)

type BusinessDetailDto struct {
	Name    string          `json:"name"`
	Website *string         `json:"website,omitempty"`
	Logo    *image.ImageDto `json:"logo,omitempty"`
}

func (r *BusinessDetailDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewBusinessDetailDto(businessDetail db.BusinessDetail) *BusinessDetailDto {
	return &BusinessDetailDto{
		Name:    businessDetail.Name,
		Website: util.NilString(businessDetail.Website),
	}
}

func (r *BusinessDetailResolver) CreateBusinessDetailDto(ctx context.Context, businessDetail db.BusinessDetail) *BusinessDetailDto {
	response := NewBusinessDetailDto(businessDetail)

	if businessDetail.LogoID.Valid {
		image, err := r.ImageResolver.Repository.GetImage(ctx, businessDetail.LogoID.Int64)
		
		if err != nil {
			util.LogOnError("OCPI221", "Error retrieving image", err)
			log.Printf("OCPI222: LogoID=%#v", businessDetail.LogoID)
			return response
		}
		
		response.Logo = r.ImageResolver.CreateImageDto(ctx, image)
	}

	return response
}
