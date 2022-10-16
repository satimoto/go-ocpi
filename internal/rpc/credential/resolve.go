package credential

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	"github.com/satimoto/go-ocpi/internal/credential"
	"github.com/satimoto/go-ocpi/internal/image"
	"github.com/satimoto/go-ocpi/internal/service"
)

type RpcCredentialResolver struct {
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	CredentialResolver     *credential.CredentialResolver
	ImageResolver          *image.ImageResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *RpcCredentialResolver {
	return &RpcCredentialResolver{
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService, services),
		ImageResolver:          image.NewResolver(repositoryService),
	}
}
