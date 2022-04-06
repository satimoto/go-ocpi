package tokenauthorization

import (
	"context"

	"github.com/satimoto/go-datastore/db"
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
	Repository TokenAuthorizationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *TokenAuthorizationResolver {
	repo := TokenAuthorizationRepository(repositoryService)
	return &TokenAuthorizationResolver{repo}
}
