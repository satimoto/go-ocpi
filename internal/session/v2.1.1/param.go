package session

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
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
