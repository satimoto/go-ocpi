package credential

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
)

type CredentialRepository interface {
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
	UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error)
}

type CredentialResolver struct {
	Repository CredentialRepository
	*businessdetail.BusinessDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	repo := CredentialRepository(repositoryService)
	return &CredentialResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
	}
}
