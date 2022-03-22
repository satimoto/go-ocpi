package businessdetail

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/image"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type BusinessDetailPayload struct {
	Name    string              `json:"name"`
	Website *string             `json:"website,omitempty"`
	Logo    *image.ImagePayload `json:"logo,omitempty"`
}

func (r *BusinessDetailPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewBusinessDetailPayload(businessDetail db.BusinessDetail) *BusinessDetailPayload {
	return &BusinessDetailPayload{
		Name:    businessDetail.Name,
		Website: util.NilString(businessDetail.Website.String),
	}
}

func (r *BusinessDetailResolver) CreateBusinessDetailPayload(ctx context.Context, businessDetail db.BusinessDetail) *BusinessDetailPayload {
	response := NewBusinessDetailPayload(businessDetail)

	if businessDetail.LogoID.Valid {
		if image, err := r.ImageResolver.Repository.GetImage(ctx, businessDetail.LogoID.Int64); err == nil {
			response.Logo = r.ImageResolver.CreateImagePayload(ctx, image)
		}
	}

	return response
}
