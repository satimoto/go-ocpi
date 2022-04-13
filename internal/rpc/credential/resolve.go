package credential

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/image"
)

type RpcCredentialRepository interface{}

type RpcCredentialResolver struct {
	Repository RpcCredentialRepository
	*businessdetail.BusinessDetailResolver
	*credential.CredentialResolver
	*image.ImageResolver
}

func NewResolver(repositoryService *db.RepositoryService) *RpcCredentialResolver {
	repo := RpcCredentialRepository(repositoryService)
	return &RpcCredentialResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
	}
}
