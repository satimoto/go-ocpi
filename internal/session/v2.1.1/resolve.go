package session

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type SessionResolver struct {
	Repository                 session.SessionRepository
	OcpiRequester              *transportation.OcpiRequester
	ChargingPeriodResolver     *chargingperiod.ChargingPeriodResolver
	LocationResolver           *location.LocationResolver
	TokenResolver              *token.TokenResolver
	TokenAuthorizationResolver *tokenauthorization.TokenAuthorizationResolver
	VersionDetailResolver      *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SessionResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *SessionResolver {
	return &SessionResolver{
		Repository:                 session.NewRepository(repositoryService),
		OcpiRequester:              ocpiRequester,
		ChargingPeriodResolver:     chargingperiod.NewResolver(repositoryService),
		LocationResolver:           location.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenResolver:              token.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
