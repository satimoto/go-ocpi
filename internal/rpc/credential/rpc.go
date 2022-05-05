package credential

import (
	"context"
	"errors"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
	ocpiCredential "github.com/satimoto/go-ocpi-api/pkg/ocpi/credential"
)

func (r *RpcCredentialResolver) CreateCredential(ctx context.Context, request *ocpirpc.CreateCredentialRequest) (*ocpirpc.CreateCredentialResponse, error) {
	if request != nil {
		params := ocpiCredential.NewCreateCredentialParams(*request)

		if request.BusinessDetail != nil {
			b, err := r.createBusinessDetail(ctx, request.BusinessDetail)

			if err != nil {
				return nil, err
			}

			params.BusinessDetailID = b.ID
		}

		c, err := r.CredentialResolver.Repository.CreateCredential(ctx, params)

		if err != nil {
			return nil, err
		}

		return r.CreateCredentialResponse(ctx, c), nil
	}

	return nil, errors.New("Missing request")
}

func (r *RpcCredentialResolver) RegisterCredential(ctx context.Context, request *ocpirpc.RegisterCredentialRequest) (*ocpirpc.RegisterCredentialResponse, error) {
	return nil, errors.New("Missing request")
}

func (r *RpcCredentialResolver) UnregisterCredential(ctx context.Context, request *ocpirpc.UnregisterCredentialRequest) (*ocpirpc.UnregisterCredentialResponse, error) {
	return nil, errors.New("Missing request")
}

func (r *RpcCredentialResolver) createBusinessDetail(ctx context.Context, request *ocpirpc.CreateBusinessDetailRequest) (*db.BusinessDetail, error) {
	if request != nil {
		params := ocpi.NewCreateBusinessDetailParams(*request)

		if request.Logo != nil {
			i, err := r.createImage(ctx, request.Logo)

			if err != nil {
				return nil, err
			}

			params.LogoID = util.SqlNullInt64(i.ID)
		}

		if b, err := r.BusinessDetailResolver.Repository.CreateBusinessDetail(ctx, params); err == nil {
			return &b, nil
		}
	}

	return nil, errors.New("Error creating business detail from RPC")
}

func (r *RpcCredentialResolver) createImage(ctx context.Context, request *ocpirpc.CreateImageRequest) (*db.Image, error) {
	if request != nil {
		params := ocpi.NewCreateImageParams(*request)

		if i, err := r.ImageResolver.Repository.CreateImage(ctx, params); err == nil {
			return &i, nil
		}
	}

	return nil, errors.New("Error creating logo from RPC")
}
