package session

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
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
	Repository SessionRepository
	*chargingperiod.ChargingPeriodResolver
	*location.LocationResolver
	*token.TokenResolver
	*tokenauthorization.TokenAuthorizationResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SessionResolver {
	repo := SessionRepository(repositoryService)
	return &SessionResolver{
		Repository:                 repo,
		ChargingPeriodResolver:     chargingperiod.NewResolver(repositoryService),
		LocationResolver:           location.NewResolver(repositoryService),
		TokenResolver:              token.NewResolver(repositoryService),
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolver(repositoryService),
	}
}
