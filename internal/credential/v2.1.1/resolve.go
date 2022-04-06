package credential

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	sync "github.com/satimoto/go-ocpi-api/internal/sync/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/version"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

type CredentialRepository interface {
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
	UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error)
}

type CredentialResolver struct {
	Repository CredentialRepository
	*businessdetail.BusinessDetailResolver
	*sync.SyncResolver
	*version.VersionResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	repo := CredentialRepository(repositoryService)
	return &CredentialResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		SyncResolver:           sync.NewResolver(repositoryService),
		VersionResolver:        version.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
