package session

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, arg db.CreateSessionParams) (db.Session, error)
	DeleteSessionChargingPeriods(ctx context.Context, sessionID int64) error
	GetSessionByUid(ctx context.Context, uid string) (db.Session, error)
	ListSessionChargingPeriods(ctx context.Context, sessionID int64) ([]db.ChargingPeriod, error)
	SetSessionChargingPeriod(ctx context.Context, arg db.SetSessionChargingPeriodParams) error
	UpdateSessionByUid(ctx context.Context, arg db.UpdateSessionByUidParams) (db.Session, error)
}

type SessionResolver struct {
	Repository SessionRepository
	*chargingperiod.ChargingPeriodResolver
	*credential.CredentialResolver
	*location.LocationResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SessionResolver {
	repo := SessionRepository(repositoryService)
	return &SessionResolver{
		Repository:             repo,
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService),
	}
}

func (r *SessionResolver) ReplaceSession(ctx context.Context, uid string, payload *SessionPayload) *db.Session {
	if payload != nil {
		session, err := r.Repository.GetSessionByUid(ctx, uid)
		locationID := util.NilInt64(session.LocationID)

		if payload.Location != nil {
			if location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *payload.Location.ID); err == nil {
				locationID = &location.ID
			} else {
				location := r.LocationResolver.ReplaceLocation(ctx, *payload.Location.ID, payload.Location)
				locationID = &location.ID
			}
		}

		if err == nil {
			sessionParams := NewUpdateSessionByUidParams(session)
			sessionParams.LocationID = *locationID

			if payload.AuthID != nil {
				sessionParams.AuthID = *payload.AuthID
			}

			if payload.AuthMethod != nil {
				sessionParams.AuthMethod = *payload.AuthMethod
			}

			if payload.Currency != nil {
				sessionParams.Currency = *payload.Currency
			}

			if payload.EndDatetime != nil {
				sessionParams.EndDatetime = util.SqlNullTime(payload.EndDatetime)
			}

			if payload.Kwh != nil {
				sessionParams.Kwh = *payload.Kwh
			}

			if payload.LastUpdated != nil {
				sessionParams.LastUpdated = *payload.LastUpdated
			}

			if payload.MeterID != nil {
				sessionParams.MeterID = util.SqlNullString(payload.MeterID)
			}
			
			if payload.Status != nil {
				sessionParams.Status = *payload.Status
			}

			if payload.StartDatetime != nil {
				sessionParams.StartDatetime = *payload.StartDatetime
			}

			if payload.TotalCost != nil {
				sessionParams.TotalCost = util.SqlNullFloat64(payload.TotalCost)
			}

			session, err = r.Repository.UpdateSessionByUid(ctx, sessionParams)
		} else {
			sessionParams := NewCreateSessionParams(payload)
			sessionParams.LocationID = *locationID

			session, err = r.Repository.CreateSession(ctx, sessionParams)
		}

		if payload.ChargingPeriods != nil {
			r.replaceChargingPeriods(ctx, session.ID, payload)
		}

		return &session
	}

	return nil
}

func (r *SessionResolver) replaceChargingPeriods(ctx context.Context, sessionID int64, payload *SessionPayload) {
	r.Repository.DeleteSessionChargingPeriods(ctx, sessionID)

	for _, chargingPeriodPayload := range payload.ChargingPeriods {
		chargingPeriod := r.ChargingPeriodResolver.ReplaceChargingPeriod(ctx, chargingPeriodPayload)

		if chargingPeriod != nil {
			r.Repository.SetSessionChargingPeriod(ctx, db.SetSessionChargingPeriodParams{
				SessionID: sessionID,
				ChargingPeriodID: chargingPeriod.ID,
			})
		}
	}
}
