package credential

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/sync"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/version"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type CredentialRepository interface {
	CreateCredential(ctx context.Context, arg db.CreateCredentialParams) (db.Credential, error)
	GetCredential(ctx context.Context, id int64) (db.Credential, error)
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
	ListCredentials(ctx context.Context) ([]db.Credential, error)
	UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error)
}

type CredentialResolver struct {
	Repository             CredentialRepository
	OcpiRequester          *transportation.OcpiRequester
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	SyncResolver           *sync.SyncResolver
	VersionResolver        *version.VersionResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *CredentialResolver {
	repo := CredentialRepository(repositoryService)

	return &CredentialResolver{
		Repository:             repo,
		OcpiRequester:          ocpiRequester,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		SyncResolver:           sync.NewResolver(repositoryService),
		VersionResolver:        version.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
