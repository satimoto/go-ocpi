package tokenauthorization

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/async"
	"github.com/satimoto/go-ocpi/internal/service"
)

type RpcTokenAuthorizationRepository interface{}

type RpcTokenAuthorizationResolver struct {
	Repository   RpcTokenAuthorizationRepository
	AsyncService *async.AsyncService
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *RpcTokenAuthorizationResolver {
	repo := RpcTokenAuthorizationRepository(repositoryService)

	return &RpcTokenAuthorizationResolver{
		Repository:   repo,
		AsyncService: services.AsyncService,
	}
}
