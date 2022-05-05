package credential

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
	ocpiCredential "github.com/satimoto/go-ocpi-api/pkg/ocpi/credential"
)

func (r *RpcCredentialResolver) CreateCredentialResponse(ctx context.Context, credential db.Credential) *ocpirpc.CreateCredentialResponse {
	response := ocpiCredential.NewCredentialResponse(credential)

	if b, err := r.BusinessDetailResolver.Repository.GetBusinessDetail(ctx, credential.BusinessDetailID); err == nil {
		response.BusinessDetail = r.createBusinessDetailResponse(ctx, b)
	}

	return response
}

func (r *RpcCredentialResolver) createBusinessDetailResponse(ctx context.Context, businessDetail db.BusinessDetail) *ocpirpc.BusinessDetailResponse {
	response := ocpi.NewBusinessDetailResponse(businessDetail)

	if businessDetail.LogoID.Valid {
		if i, err := r.ImageResolver.Repository.GetImage(ctx, businessDetail.LogoID.Int64); err == nil {
			response.Logo = ocpi.NewImageResponse(i)
		}
	}

	return response
}
