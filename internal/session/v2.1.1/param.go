package session

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func NewCreateSessionParams(sessionDto *dto.SessionDto) db.CreateSessionParams {
	return db.CreateSessionParams{
		Uid:             *sessionDto.ID,
		AuthorizationID: util.SqlNullString(sessionDto.AuthorizationID),
		StartDatetime:   *sessionDto.StartDatetime,
		EndDatetime:     util.SqlNullTime(sessionDto.EndDatetime),
		Kwh:             *sessionDto.Kwh,
		AuthID:          *sessionDto.AuthID,
		AuthMethod:      *sessionDto.AuthMethod,
		MeterID:         util.SqlNullString(sessionDto.MeterID),
		Currency:        *sessionDto.Currency,
		TotalCost:       util.SqlNullFloat64(sessionDto.TotalCost),
		Status:          *sessionDto.Status,
		LastUpdated:     *sessionDto.LastUpdated,
	}
}
