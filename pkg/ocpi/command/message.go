package command

import (
	"encoding/hex"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewCommandReservationResponse(command db.CommandReservation) *ocpirpc.ReserveNowResponse {
	return &ocpirpc.ReserveNowResponse{
		Id:            command.ID,
		Status:        string(command.Status),
		ReservationId: command.ReservationID,
		ExpiryDate:    command.ExpiryDate.Format(time.RFC3339),
		LocationUid:   command.LocationID,
		EvseUid:       command.EvseUid.String,
	}
}

func NewCommandStartResponse(command db.CommandStart, verificationKey []byte) *ocpirpc.StartSessionResponse {
	return &ocpirpc.StartSessionResponse{
		Id:              command.ID,
		Status:          string(command.Status),
		AuthorizationId: command.AuthorizationID.String,
		VerificationKey: hex.EncodeToString(verificationKey),
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
