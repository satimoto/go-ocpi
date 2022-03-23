package versiondetail

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type VersionDetailRepository interface {
	CreateVersionEndpoint(ctx context.Context, arg db.CreateVersionEndpointParams) (db.VersionEndpoint, error)
	DeleteVersionEndpoints(ctx context.Context, versionID int64) error
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
	GetVersionEndpoint(ctx context.Context, id int64) (db.VersionEndpoint, error)
	ListVersionEndpoints(ctx context.Context, versionID int64) ([]db.VersionEndpoint, error)
	UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error)
}

type VersionDetailResolver struct {
	Repository VersionDetailRepository
	*util.OCPIRequester
}

func NewResolver(repositoryService *db.RepositoryService) *VersionDetailResolver {
	repo := VersionDetailRepository(repositoryService)

	return &VersionDetailResolver{
		Repository:    repo,
		OCPIRequester: util.NewOCPIRequester(),
	}
}

func (r *VersionDetailResolver) ReplaceVersionEndpoints(ctx context.Context, versionID int64, payload *VersionDetailPayload) []*db.VersionEndpoint {
	endpoints := []*db.VersionEndpoint{}

	if payload != nil {
		r.Repository.DeleteVersionEndpoints(ctx, versionID)

		for _, endpointPayload := range payload.Endpoints {
			endpointParams := NewCreateVersionEndpointParams(versionID, endpointPayload)

			if endpoint, err := r.Repository.CreateVersionEndpoint(ctx, endpointParams); err == nil {
				endpoints = append(endpoints, &endpoint)
			}
		}

	}

	return endpoints
}

func (r *VersionDetailResolver) UpdateVersionDetail(ctx context.Context, url string, header util.OCPIRequestHeader, versionID int64) []*db.VersionEndpoint {
	if response, err := r.OCPIRequester.Do("GET", url, header, nil); err == nil {
		defer response.Close()

		ocpiResponse, err := r.UnmarshalResponse(response)

		if err == nil && ocpiResponse.StatusCode == ocpi.STATUS_CODE_OK {
			return r.ReplaceVersionEndpoints(ctx, versionID, ocpiResponse.Data)
		}
	}

	return []*db.VersionEndpoint{}
}
