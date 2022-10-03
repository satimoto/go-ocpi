package credential

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	"github.com/satimoto/go-ocpi/internal/credential"
	"github.com/satimoto/go-ocpi/internal/image"
	"github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type RpcCredentialResolver struct {
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	CredentialResolver     *credential.CredentialResolver
	ImageResolver          *image.ImageResolver
}

func NewResolver(repositoryService *db.RepositoryService, syncService *sync.SyncService, ocpiRequester *transportation.OcpiRequester) *RpcCredentialResolver {
	return &RpcCredentialResolver{
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService, syncService, ocpiRequester),
		ImageResolver:          image.NewResolver(repositoryService),
	}
}
