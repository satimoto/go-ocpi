package credential

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"github.com/satimoto/go-ocpi-api/ocpirpc/credentialrpc"
)

func NewBusinessDetailIto(businessDetail db.BusinessDetail) *ocpirpc.BusinessDetail {
	return &ocpirpc.BusinessDetail{
		Name:    businessDetail.Name,
		Website: businessDetail.Website.String,
	}
}

func NewCredentialIto(credential db.Credential) *credentialrpc.CreateCredentialResponse {
	return &credentialrpc.CreateCredentialResponse{
		Id:          credential.ID,
		ClientToken: credential.ClientToken.String,
		Url:         credential.Url,
		PartyId:     credential.PartyID,
		CountryCode: credential.CountryCode,
		IsHub:       credential.IsHub,
	}
}

func NewImageIto(image db.Image) *ocpirpc.Image {
	return &ocpirpc.Image{
		Url:       image.Url,
		Thumbnail: image.Thumbnail.String,
		Category:  string(image.Category),
		Type:      image.Type,
		Width:     image.Width.Int32,
		Height:    image.Height.Int32,
	}
}

func (r *RpcCredentialResolver) CreateCredentialIto(ctx context.Context, credential db.Credential) *credentialrpc.CreateCredentialResponse {
	response := NewCredentialIto(credential)

	if b, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, credential.BusinessDetailID); err == nil {
		response.BusinessDetail = r.createBusinessDetailIto(ctx, b)
	}

	return response
}

func (r *RpcCredentialResolver) createBusinessDetailIto(ctx context.Context, businessDetail db.BusinessDetail) *ocpirpc.BusinessDetail {
	response := NewBusinessDetailIto(businessDetail)

	if businessDetail.LogoID.Valid {
		if i, err := r.ImageResolver.Repository.GetImage(ctx, businessDetail.LogoID.Int64); err == nil {
			response.Logo = NewImageIto(i)
		}
	}

	return response
}
