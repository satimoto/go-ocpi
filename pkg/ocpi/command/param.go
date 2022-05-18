package command

import (
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateCommandReservationParams(token db.Token, expiryDate time.Time, location db.Location, evseUid *string) db.CreateCommandReservationParams {
	return db.CreateCommandReservationParams{
		Status:     db.CommandResponseTypeREQUESTED,
		TokenID:    token.ID,
		ExpiryDate: expiryDate,
		LocationID: location.Uid,
		EvseUid:    util.SqlNullString(evseUid),
	}
}

func NewCreateCommandStartParams(token db.Token, location db.Location, evseUid *string) db.CreateCommandStartParams {
	return db.CreateCommandStartParams{
		Status:     db.CommandResponseTypeREQUESTED,
		TokenID:    token.ID,
		LocationID: location.Uid,
		EvseUid:    util.SqlNullString(evseUid),
	}
}

func NewCreateCommandStopParams(sessionID string) db.CreateCommandStopParams {
	return db.CreateCommandStopParams{
		Status:    db.CommandResponseTypeREQUESTED,
		SessionID: sessionID,
	}
}

func NewCreateCommandUnlockParams(location db.Location, evseUid string, connectorID string) db.CreateCommandUnlockParams {
	return db.CreateCommandUnlockParams{
		Status:      db.CommandResponseTypeREQUESTED,
		LocationID:  location.Uid,
		EvseUid:     evseUid,
		ConnectorID: connectorID,
	}
}
