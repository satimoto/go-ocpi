package tokenauthorization

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
)

type TokenAuthorizationResolver struct {
	Repository        tokenauthorization.TokenAuthorizationRepository
	ConnectorResolver *connector.ConnectorResolver
	EvseResolver      *evse.EvseResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TokenAuthorizationResolver {
	return &TokenAuthorizationResolver{
		Repository:        tokenauthorization.NewRepository(repositoryService),
		ConnectorResolver: connector.NewResolver(repositoryService),
		EvseResolver:      evse.NewResolver(repositoryService),
	}
}
