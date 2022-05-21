package credential

import (
	credentialRepository "github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/version"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type CredentialResolver struct {
	Repository             credentialRepository.CredentialRepository
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	CredentialResolver     *credential.CredentialResolver
	VersionResolver        *version.VersionResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	return &CredentialResolver{
		Repository:             credentialRepository.NewRepository(repositoryService),
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService),
		VersionResolver:        version.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
