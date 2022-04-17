package commandrpc

import (
	"time"

	"github.com/satimoto/go-datastore/db"
)

func NewCommandReservationResponse(command db.CommandReservation) *ReserveNowResponse {
	return &ReserveNowResponse{
		Id:            command.ID,
		Status:        string(command.Status),
		ReservationId: command.ReservationID,
		ExpiryDate:    command.ExpiryDate.Format(time.RFC3339Nano),
		LocationUid:   command.LocationID,
		EvseUid:       command.EvseUid.String,
	}
}

func NewCommandStartResponse(command db.CommandStart) *StartSessionResponse {
	return &StartSessionResponse{
		Id:              command.ID,
		Status:          string(command.Status),
		AuthorizationId: command.AuthorizationID.String,
		LocationUid:     command.LocationID,
		EvseUid:         command.EvseUid.String,
	}
}

func NewCommandStopResponse(command db.CommandStop) *StopSessionResponse {
	return &StopSessionResponse{
		Id:         command.ID,
		Status:     string(command.Status),
		SessionUid: command.SessionID,
	}
}

func NewCommandUnlockResponse(command db.CommandUnlock) *UnlockConnectorResponse {
	return &UnlockConnectorResponse{
		Id:           command.ID,
		Status:       string(command.Status),
		LocationUid:  command.LocationID,
		EvseUid:      command.EvseUid,
		ConnectorUid: command.ConnectorID,
	}
}
