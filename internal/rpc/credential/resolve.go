package credential

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	"github.com/satimoto/go-ocpi-api/internal/image"
)

type RpcCredentialResolver struct {
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	CredentialResolver     *credential.CredentialResolver
	ImageResolver          *image.ImageResolver
}

func NewResolver(repositoryService *db.RepositoryService) *RpcCredentialResolver {
	return &RpcCredentialResolver{
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
	}
}
