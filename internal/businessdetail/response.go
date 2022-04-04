package businessdetail

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/image"
	"github.com/satimoto/go-ocpi-api/internal/util"
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
		Website: util.NilString(businessDetail.Website.String),
	}
}

func (r *BusinessDetailResolver) CreateBusinessDetailDto(ctx context.Context, businessDetail db.BusinessDetail) *BusinessDetailDto {
	response := NewBusinessDetailDto(businessDetail)

	if businessDetail.LogoID.Valid {
		if image, err := r.ImageResolver.Repository.GetImage(ctx, businessDetail.LogoID.Int64); err == nil {
			response.Logo = r.ImageResolver.CreateImageDto(ctx, image)
		}
	}

	return response
}
