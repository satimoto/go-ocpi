package session

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/tokenauthorization"
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
	*location.LocationResolver
	*tokenauthorization.TokenAuthorizationResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SessionResolver {
	repo := SessionRepository(repositoryService)
	return &SessionResolver{
		Repository:             repo,
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService),
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
	}
}

func (r *SessionResolver) ReplaceSession(ctx context.Context, countryCode *string, partyID *string, uid string, dto *SessionDto) *db.Session {
	if dto != nil {
		session, err := r.Repository.GetSessionByUid(ctx, uid)
		locationID := util.NilInt64(session.LocationID)

		if dto.Location != nil {
			if location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *dto.Location.ID); err == nil {
				locationID = &location.ID
			} else {
				location := r.LocationResolver.ReplaceLocationWithIdentifiers(ctx, countryCode, partyID, *dto.Location.ID, dto.Location)
				locationID = &location.ID
			}
		}

		if err == nil {
			sessionParams := NewUpdateSessionByUidParams(session)
			sessionParams.CountryCode = util.SqlNullString(countryCode)
			sessionParams.PartyID = util.SqlNullString(partyID)
			sessionParams.LocationID = *locationID

			if dto.AuthID != nil {
				sessionParams.AuthID = *dto.AuthID
			}

			if dto.AuthMethod != nil {
				sessionParams.AuthMethod = *dto.AuthMethod
			}

			if dto.Currency != nil {
				sessionParams.Currency = *dto.Currency
			}

			if dto.EndDatetime != nil {
				sessionParams.EndDatetime = util.SqlNullTime(dto.EndDatetime)
			}

			if dto.Kwh != nil {
				sessionParams.Kwh = *dto.Kwh
			}

			if dto.LastUpdated != nil {
				sessionParams.LastUpdated = *dto.LastUpdated
			}

			if dto.MeterID != nil {
				sessionParams.MeterID = util.SqlNullString(dto.MeterID)
			}

			if dto.Status != nil {
				sessionParams.Status = *dto.Status
			}

			if dto.StartDatetime != nil {
				sessionParams.StartDatetime = *dto.StartDatetime
			}

			if dto.TotalCost != nil {
				sessionParams.TotalCost = util.SqlNullFloat64(dto.TotalCost)
			}

			session, err = r.Repository.UpdateSessionByUid(ctx, sessionParams)
		} else {
			sessionParams := NewCreateSessionParams(dto)
			sessionParams.CountryCode = util.SqlNullString(countryCode)
			sessionParams.PartyID = util.SqlNullString(partyID)
			sessionParams.LocationID = *locationID

			session, err = r.Repository.CreateSession(ctx, sessionParams)
		}

		if dto.AuthorizationID != nil {
			r.replaceTokenAuthorization(ctx, countryCode, partyID, dto)
		}

		if dto.ChargingPeriods != nil {
			r.replaceChargingPeriods(ctx, session.ID, dto)
		}

		return &session
	}

	return nil
}

func (r *SessionResolver) replaceChargingPeriods(ctx context.Context, sessionID int64, dto *SessionDto) {
	r.Repository.DeleteSessionChargingPeriods(ctx, sessionID)

	for _, chargingPeriodDto := range dto.ChargingPeriods {
		chargingPeriod := r.ChargingPeriodResolver.ReplaceChargingPeriod(ctx, chargingPeriodDto)

		if chargingPeriod != nil {
			r.Repository.SetSessionChargingPeriod(ctx, db.SetSessionChargingPeriodParams{
				SessionID:        sessionID,
				ChargingPeriodID: chargingPeriod.ID,
			})
		}
	}
}

func (r *SessionResolver) replaceTokenAuthorization(ctx context.Context, countryCode *string, partyID *string, dto *SessionDto) {
	tokenAuthorizationParams := tokenauthorization.NewUpdateTokenAuthorizationParams(*dto.AuthorizationID, countryCode, partyID)
	r.TokenAuthorizationResolver.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, tokenAuthorizationParams)
}
