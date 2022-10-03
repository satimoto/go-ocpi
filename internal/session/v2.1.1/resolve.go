package session

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/token"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	"github.com/satimoto/go-ocpi/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type SessionResolver struct {
	Repository                   session.SessionRepository
	OcpiRequester                *transportation.OcpiRequester
	ChargingPeriodResolver       *chargingperiod.ChargingPeriodResolver
	LocationResolver             *location.LocationResolver
	NodeRepository               node.NodeRepository
	TokenRepository              token.TokenRepository
	TokenAuthorizationRepository tokenauthorization.TokenAuthorizationRepository
	VersionDetailResolver        *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SessionResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *SessionResolver {
	return &SessionResolver{
		Repository:                   session.NewRepository(repositoryService),
		OcpiRequester:                ocpiRequester,
		ChargingPeriodResolver:       chargingperiod.NewResolver(repositoryService),
		LocationResolver:             location.NewResolverWithServices(repositoryService, ocpiRequester),
		NodeRepository:               node.NewRepository(repositoryService),
		TokenRepository:              token.NewRepository(repositoryService),
		TokenAuthorizationRepository: tokenauthorization.NewRepository(repositoryService),
		VersionDetailResolver:        versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
