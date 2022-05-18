package command

import (
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewCommandReservationResponse(command db.CommandReservation) *ocpirpc.ReserveNowResponse {
	return &ocpirpc.ReserveNowResponse{
		Id:            command.ID,
		Status:        string(command.Status),
		ReservationId: command.ReservationID,
		ExpiryDate:    command.ExpiryDate.Format(time.RFC3339Nano),
		LocationUid:   command.LocationID,
		EvseUid:       command.EvseUid.String,
	}
}

func NewCommandStartResponse(command db.CommandStart) *ocpirpc.StartSessionResponse {
	return &ocpirpc.StartSessionResponse{
		Id:              command.ID,
		Status:          string(command.Status),
		AuthorizationId: command.AuthorizationID.String,
		LocationUid:     command.LocationID,
		EvseUid:         command.EvseUid.String,
	}
}

func NewCommandStopResponse(command db.CommandStop) *ocpirpc.StopSessionResponse {
	return &ocpirpc.StopSessionResponse{
		Id:         command.ID,
		Status:     string(command.Status),
		SessionUid: command.SessionID,
	}
}

func NewCommandUnlockResponse(command db.CommandUnlock) *ocpirpc.UnlockConnectorResponse {
	return &ocpirpc.UnlockConnectorResponse{
		Id:           command.ID,
		Status:       string(command.Status),
		LocationUid:  command.LocationID,
		EvseUid:      command.EvseUid,
		ConnectorUid: command.ConnectorID,
	}
}
