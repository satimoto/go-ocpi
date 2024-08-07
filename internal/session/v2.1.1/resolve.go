package session

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/token"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	"github.com/satimoto/go-ocpi/internal/async"
	"github.com/satimoto/go-ocpi/internal/chargingperiod"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/service"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type SessionResolver struct {
	Repository                   session.SessionRepository
	OcpiService                  *transportation.OcpiService
	AsyncService                 *async.AsyncService
	ChargingPeriodResolver       *chargingperiod.ChargingPeriodResolver
	CommandResolver              *command.CommandResolver
	LocationResolver             *location.LocationResolver
	NodeRepository               node.NodeRepository
	TokenRepository              token.TokenRepository
	TokenAuthorizationRepository tokenauthorization.TokenAuthorizationRepository
	VersionDetailResolver        *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *SessionResolver {
	return &SessionResolver{
		Repository:                   session.NewRepository(repositoryService),
		OcpiService:                  services.OcpiService,
		AsyncService:                 services.AsyncService,
		ChargingPeriodResolver:       chargingperiod.NewResolver(repositoryService),
		CommandResolver:              command.NewResolver(repositoryService, services),
		LocationResolver:             location.NewResolver(repositoryService, services),
		NodeRepository:               node.NewRepository(repositoryService),
		TokenRepository:              token.NewRepository(repositoryService),
		TokenAuthorizationRepository: tokenauthorization.NewRepository(repositoryService),
		VersionDetailResolver:        versiondetail.NewResolver(repositoryService, services),
	}
}
