package credential

import (
	"context"
	"errors"
	"log"

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
			businessDetail, err := r.createBusinessDetail(ctx, request.BusinessDetail)

			if err != nil {
				return nil, err
			}

			params.BusinessDetailID = businessDetail.ID
		}

		credential, err := r.CredentialResolver.Repository.CreateCredential(ctx, params)

		if err != nil {
			util.LogOnError("OCPI002", "Error creating business detail", err)
			log.Printf("OCPI002: Params=%#v", params)
			return nil, errors.New("Error creating business detail")
		}

		return r.CreateCredentialResponse(ctx, credential), nil
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
			image, err := r.createImage(ctx, request.Logo)

			if err != nil {
				return nil, err
			}

			params.LogoID = util.SqlNullInt64(image.ID)
		}

		businessDetail, err := r.BusinessDetailResolver.Repository.CreateBusinessDetail(ctx, params)

		if err != nil {
			util.LogOnError("OCPI003", "Error creating business detail", err)
			log.Printf("OCPI003: Params=%#v", params)
			return nil, errors.New("Error creating business detail")
		}

		return &businessDetail, nil
	}

	return nil, errors.New("Error creating business detail from RPC")
}

func (r *RpcCredentialResolver) createImage(ctx context.Context, request *ocpirpc.CreateImageRequest) (*db.Image, error) {
	if request != nil {
		params := ocpi.NewCreateImageParams(*request)
		image, err := r.ImageResolver.Repository.CreateImage(ctx, params)

		if err != nil {
			util.LogOnError("OCPI004", "Error creating image", err)
			log.Printf("OCPI004: Params=%#v", params)
			return nil, errors.New("Error creating image")
		}
	
		return &image, nil
	}

	return nil, errors.New("Error creating logo from RPC")
}
