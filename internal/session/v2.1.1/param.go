package session

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateSessionParams(dto *SessionDto) db.CreateSessionParams {
	return db.CreateSessionParams{
		Uid:             *dto.ID,
		AuthorizationID: util.SqlNullString(dto.AuthorizationID),
		StartDatetime:   *dto.StartDatetime,
		EndDatetime:     util.SqlNullTime(dto.EndDatetime),
		Kwh:             *dto.Kwh,
		AuthID:          *dto.AuthID,
		AuthMethod:      *dto.AuthMethod,
		MeterID:         util.SqlNullString(dto.MeterID),
		Currency:        *dto.Currency,
		TotalCost:       util.SqlNullFloat64(dto.TotalCost),
		Status:          *dto.Status,
		LastUpdated:     *dto.LastUpdated,
	}
}

func NewUpdateSessionByUidParams(session db.Session) db.UpdateSessionByUidParams {
	return db.UpdateSessionByUidParams{
		Uid:             session.Uid,
		CountryCode:     session.CountryCode,
		PartyID:         session.PartyID,
		AuthorizationID: session.AuthorizationID,
		StartDatetime:   session.StartDatetime,
		EndDatetime:     session.EndDatetime,
		Kwh:             session.Kwh,
		AuthID:          session.AuthID,
		AuthMethod:      session.AuthMethod,
		LocationID:      session.LocationID,
		MeterID:         session.MeterID,
		Currency:        session.Currency,
		TotalCost:       session.TotalCost,
		Status:          session.Status,
		LastUpdated:     session.LastUpdated,
	}
}
