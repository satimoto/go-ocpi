package session

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error)
	DeleteSessionChargingPeriods(ctx context.Context, sessionID int64) error
	GetSessionByLastUpdated(ctx context.Context, arg db.GetSessionByLastUpdatedParams) (db.Session, error)
	GetSessionByUid(ctx context.Context, uid string) (db.Session, error)
	ListSessionChargingPeriods(ctx context.Context, sessionID int64) ([]db.ChargingPeriod, error)
	SetSessionChargingPeriod(ctx context.Context, arg db.SetSessionChargingPeriodParams) error
	UpdateSessionByUid(ctx context.Context, arg db.UpdateSessionByUidParams) (db.Session, error)
}

type SessionResolver struct {
	Repository                 SessionRepository
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
	repo := SessionRepository(repositoryService)

	return &SessionResolver{
		Repository:                 repo,
		OcpiRequester:              ocpiRequester,
		ChargingPeriodResolver:     chargingperiod.NewResolver(repositoryService),
		LocationResolver:           location.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenResolver:              token.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
