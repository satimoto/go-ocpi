package command

import (
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateCommandReservationParams(token db.Token, expiryDate time.Time, location db.Location, evseUid *string) db.CreateCommandReservationParams {
	return db.CreateCommandReservationParams{
		Status:      db.CommandResponseTypeREQUESTED,
		TokenID:     token.ID,
		ExpiryDate:  expiryDate,
		LocationID:  location.Uid,
		EvseUid:     util.SqlNullString(evseUid),
		LastUpdated: time.Now().UTC(),
	}
}

func NewCreateCommandStartParams(tokenAuthorization db.TokenAuthorization, evseUid *string) db.CreateCommandStartParams {
	createCommandStartParams := db.CreateCommandStartParams{
		AuthorizationID: util.SqlNullString(tokenAuthorization.AuthorizationID),
		Status:          db.CommandResponseTypeREQUESTED,
		TokenID:         tokenAuthorization.TokenID,
		EvseUid:         util.SqlNullString(evseUid),
		LastUpdated:     time.Now().UTC(),
	}

	if tokenAuthorization.LocationID.Valid {
		createCommandStartParams.LocationID = tokenAuthorization.LocationID.String
	}

	return createCommandStartParams
}

func NewCreateCommandStopParams(sessionID string) db.CreateCommandStopParams {
	return db.CreateCommandStopParams{
		Status:      db.CommandResponseTypeREQUESTED,
		SessionID:   sessionID,
		LastUpdated: time.Now().UTC(),
	}
}

func NewCreateCommandUnlockParams(location db.Location, evseUid string, connectorID string) db.CreateCommandUnlockParams {
	return db.CreateCommandUnlockParams{
		Status:      db.CommandResponseTypeREQUESTED,
		LocationID:  location.Uid,
		EvseUid:     evseUid,
		ConnectorID: connectorID,
		LastUpdated: time.Now().UTC(),
	}
}
