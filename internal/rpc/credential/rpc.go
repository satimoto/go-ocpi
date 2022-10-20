package credential

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	ocpiCredential "github.com/satimoto/go-ocpi/pkg/ocpi/credential"
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
			return nil, errors.New("error creating business detail")
		}

		return r.CreateCredentialResponse(ctx, credential), nil
	}

	return nil, errors.New("Missing request")
}

func (r *RpcCredentialResolver) RegisterCredential(ctx context.Context, request *ocpirpc.RegisterCredentialRequest) (*ocpirpc.RegisterCredentialResponse, error) {
	if request != nil {
		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, request.Id)

		if err != nil {
			util.LogOnError("OCPI005", "Error retrieving credential", err)
			log.Printf("OCPI005: CredentialID=%v", request.Id)
			return nil, errors.New("error registering credential")
		}

		token := credential.ClientToken.String

		if len(request.ClientToken) > 0 {
			token = request.ClientToken
		}

		_, err = r.CredentialResolver.RegisterCredential(ctx, credential, token)

		if err != nil {
			util.LogOnError("OCPI006", "Error registering credential", err)
			log.Printf("OCPI006: CredentialID=%v, Token=%v", credential.ID, token)
			return nil, errors.New("error registering credential")
		}

		return &ocpirpc.RegisterCredentialResponse{Id: credential.ID}, nil
	}

	return nil, errors.New("Missing request")
}

func (r *RpcCredentialResolver) SyncCredential(ctx context.Context, request *ocpirpc.SyncCredentialRequest) (*ocpirpc.SyncCredentialResponse, error) {
	if request != nil {
		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, request.Id)

		if err != nil {
			util.LogOnError("OCPI284", "Error retrieving credential", err)
			log.Printf("OCPI284: CredentialID=%v", request.Id)
			return nil, errors.New("error syncing credential")
		}

		fullSync := true
		lastUpdated := util.ParseTime(request.FromDate, nil)
		countryCode := util.NilString(request.CountryCode)
		partyID := util.NilString(request.PartyId)

		if lastUpdated != nil {
			fullSync = false
		}

		go r.CredentialResolver.SyncService.SynchronizeCredential(credential, fullSync, lastUpdated, countryCode, partyID)

		return &ocpirpc.SyncCredentialResponse{Id: credential.ID}, nil
	}

	return nil, errors.New("Missing request")
}

func (r *RpcCredentialResolver) UnregisterCredential(ctx context.Context, request *ocpirpc.UnregisterCredentialRequest) (*ocpirpc.UnregisterCredentialResponse, error) {
	if request != nil {
		credential, err := r.CredentialResolver.Repository.GetCredential(ctx, request.Id)

		if err != nil {
			util.LogOnError("OCPI007", "Error retrieving credential", err)
			log.Printf("OCPI007: CredentialID=%v", request.Id)
			return nil, errors.New("error unregistering credential")
		}

		_, err = r.CredentialResolver.UnregisterCredential(ctx, credential)

		if err != nil {
			util.LogOnError("OCPI008", "Error unregistering credential", err)
			log.Printf("OCPI008: CredentialID=%v", credential.ID)
			return nil, errors.New("error unregistering credential")
		}

		return &ocpirpc.UnregisterCredentialResponse{Id: credential.ID}, nil
	}

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
			return nil, errors.New("error creating business detail")
		}

		return &businessDetail, nil
	}

	return nil, errors.New("error creating business detail from RPC")
}

func (r *RpcCredentialResolver) createImage(ctx context.Context, request *ocpirpc.CreateImageRequest) (*db.Image, error) {
	if request != nil {
		params := ocpi.NewCreateImageParams(*request)
		image, err := r.ImageResolver.Repository.CreateImage(ctx, params)

		if err != nil {
			util.LogOnError("OCPI004", "Error creating image", err)
			log.Printf("OCPI004: Params=%#v", params)
			return nil, errors.New("error creating image")
		}

		return &image, nil
	}

	return nil, errors.New("error creating logo from RPC")
}
