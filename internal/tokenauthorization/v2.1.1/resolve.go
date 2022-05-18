package tokenauthorization

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
)

type TokenAuthorizationRepository interface {
	CreateTokenAuthorization(ctx context.Context, arg db.CreateTokenAuthorizationParams) (db.TokenAuthorization, error)
	GetTokenAuthorizationByAuthorizationID(ctx context.Context, authorizationID string) (db.TokenAuthorization, error)
	ListTokenAuthorizationConnectors(ctx context.Context, tokenAuthorizationID int64) ([]db.Connector, error)
	ListTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64) ([]db.Evse, error)
	SetTokenAuthorizationConnector(ctx context.Context, arg db.SetTokenAuthorizationConnectorParams) error
	SetTokenAuthorizationEvse(ctx context.Context, arg db.SetTokenAuthorizationEvseParams) error
	UpdateTokenAuthorizationByAuthorizationID(ctx context.Context, arg db.UpdateTokenAuthorizationByAuthorizationIDParams) (db.TokenAuthorization, error)
}

type TokenAuthorizationResolver struct {
	Repository        TokenAuthorizationRepository
	ConnectorResolver *connector.ConnectorResolver
	EvseResolver      *evse.EvseResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TokenAuthorizationResolver {
	repo := TokenAuthorizationRepository(repositoryService)

	return &TokenAuthorizationResolver{
		Repository:        repo,
		ConnectorResolver: connector.NewResolver(repositoryService),
		EvseResolver:      evse.NewResolver(repositoryService),
	}
}
