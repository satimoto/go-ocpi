package version

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/util"
)

type VersionRepository interface {
	CreateVersion(ctx context.Context, arg db.CreateVersionParams) (db.Version, error)
	DeleteVersions(ctx context.Context, credentialID int64) error
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
	GetVersion(ctx context.Context, id int64) (db.Version, error)
	ListVersions(ctx context.Context, credentialID int64) ([]db.Version, error)
	UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error)
}

type VersionResolver struct {
	Repository VersionRepository
	*util.OCPIRequester
}

func NewResolver(repositoryService *db.RepositoryService) *VersionResolver {
	repo := VersionRepository(repositoryService)

	return &VersionResolver{
		Repository:         repo,
		OCPIRequester:      util.NewOCPIRequester(),
	}
}

func (r *VersionResolver) PickBestVersion(ctx context.Context, credentialID int64) *db.Version {
	if versions, err := r.Repository.ListVersions(ctx, credentialID); err == nil {
		for _, version := range versions {
			if version.Version == "2.1.1" {
				return &version
			}
		}
	}

	return nil
}

func (r *VersionResolver) ReplaceVersions(ctx context.Context, credentialID int64, payload []*VersionPayload) []*db.Version {
	versions := []*db.Version{}

	if payload != nil {
		r.Repository.DeleteVersions(ctx, credentialID)

		for _, versionPayload := range payload {
			versionParams := NewCreateVersionParams(credentialID, versionPayload)

			if version, err := r.Repository.CreateVersion(ctx, versionParams); err == nil {
				versions = append(versions, &version)
			}
		}

	}

	return versions
}
