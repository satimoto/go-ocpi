package tokenauthorization

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/token"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	"github.com/satimoto/go-ocpi/internal/async"
	"github.com/satimoto/go-ocpi/internal/service"
)

type RpcTokenAuthorizationRepository interface{}

type RpcTokenAuthorizationResolver struct {
	Repository                   RpcTokenAuthorizationRepository
	AsyncService                 *async.AsyncService
	SessionRepository            session.SessionRepository
	TokenRepository              token.TokenRepository
	TokenAuthorizationRepository tokenauthorization.TokenAuthorizationRepository
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *RpcTokenAuthorizationResolver {
	repo := RpcTokenAuthorizationRepository(repositoryService)

	return &RpcTokenAuthorizationResolver{
		Repository:                   repo,
		AsyncService:                 services.AsyncService,
		SessionRepository:            session.NewRepository(repositoryService),
		TokenRepository:              token.NewRepository(repositoryService),
		TokenAuthorizationRepository: tokenauthorization.NewRepository(repositoryService),
	}
}
