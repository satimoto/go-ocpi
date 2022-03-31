package command

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/satimoto/go-datastore/db"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type CommandReservationPayload struct {
	ResponseUrl   *string             `json:"response_url"`
	Token         *token.TokenPayload `json:"token"`
	ExpiryDate    *time.Time          `json:"expiry_date"`
	ReservationID *int64              `json:"reservation_id"`
	LocationID    *string             `json:"location_id"`
	EvseUid       *string             `json:"evse_uid,omitempty"`
}

func (r *CommandReservationPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandReservationPayload(command db.CommandReservation) *CommandReservationPayload {
	responseUrl := fmt.Sprintf("%s/%s/commands/reserve/%v", os.Getenv("API_DOMAIN"), API_VERSION, command.ID)

	return &CommandReservationPayload{
		ResponseUrl:   &responseUrl,
		ExpiryDate:    &command.ExpiryDate,
		ReservationID: &command.ReservationID,
		LocationID:    &command.LocationID,
		EvseUid:       util.NilString(command.EvseUid.String),
	}
}

func NewUpdateCommandReservationParams(command db.CommandReservation) db.UpdateCommandReservationParams {
	return db.UpdateCommandReservationParams{
		ID:         command.ID,
		Status:     command.Status,
		ExpiryDate: command.ExpiryDate,
		EvseUid:    command.EvseUid,
	}
}

func (r *CommandResolver) CreateCommandReservationPayload(ctx context.Context, command db.CommandReservation) *CommandReservationPayload {
	response := NewCommandReservationPayload(command)

	if token, err := r.TokenResolver.Repository.GetToken(ctx, command.TokenID); err == nil {
		response.Token = r.TokenResolver.CreateTokenPayload(ctx, token)
	}

	return response
}

type CommandResponsePayload struct {
	Result *db.CommandResponseType `json:"result"`
}

func (r *CommandResponsePayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type CommandStartPayload struct {
	ResponseUrl *string             `json:"response_url"`
	Token       *token.TokenPayload `json:"token"`
	LocationID  *string             `json:"location_id"`
	EvseUid     *string             `json:"evse_uid,omitempty"`
}

func (r *CommandStartPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandStartPayload(command db.CommandStart) *CommandStartPayload {
	responseUrl := fmt.Sprintf("%s/%s/commands/start/%v", os.Getenv("API_DOMAIN"), API_VERSION, command.ID)

	return &CommandStartPayload{
		ResponseUrl: &responseUrl,
		LocationID:  &command.LocationID,
		EvseUid:     util.NilString(command.EvseUid.String),
	}
}

func NewUpdateCommandStartParams(command db.CommandStart) db.UpdateCommandStartParams {
	return db.UpdateCommandStartParams{
		ID:     command.ID,
		Status: command.Status,
	}
}

func (r *CommandResolver) CreateCommandStartPayload(ctx context.Context, command db.CommandStart) *CommandStartPayload {
	response := NewCommandStartPayload(command)

	if token, err := r.TokenResolver.Repository.GetToken(ctx, command.TokenID); err == nil {
		response.Token = r.TokenResolver.CreateTokenPayload(ctx, token)
	}

	return response
}

type CommandStopPayload struct {
	ResponseUrl *string `json:"response_url"`
	SessionID   *string `json:"session_id"`
}

func (r *CommandStopPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandStopPayload(command db.CommandStop) *CommandStopPayload {
	responseUrl := fmt.Sprintf("%s/%s/commands/stop/%v", os.Getenv("API_DOMAIN"), API_VERSION, command.ID)

	return &CommandStopPayload{
		ResponseUrl: &responseUrl,
		SessionID:   &command.SessionID,
	}
}

func NewUpdateCommandStopParams(command db.CommandStop) db.UpdateCommandStopParams {
	return db.UpdateCommandStopParams{
		ID:     command.ID,
		Status: command.Status,
	}
}

func (r *CommandResolver) CreateCommandStopPayload(ctx context.Context, command db.CommandStop) *CommandStopPayload {
	return NewCommandStopPayload(command)
}

type CommandUnlockPayload struct {
	ResponseUrl *string `json:"response_url"`
	LocationID  *string `json:"location_id"`
	EvseUid     *string `json:"evse_uid"`
	ConnectorID *string `json:"connector_id"`
}

func (r *CommandUnlockPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandUnlockPayload(command db.CommandUnlock) *CommandUnlockPayload {
	responseUrl := fmt.Sprintf("%s/%s/commands/unlock/%v", os.Getenv("API_DOMAIN"), API_VERSION, command.ID)

	return &CommandUnlockPayload{
		ResponseUrl: &responseUrl,
		LocationID:  &command.LocationID,
		EvseUid:     &command.EvseUid,
		ConnectorID: &command.ConnectorID,
	}
}

func NewUpdateCommandUnlockParams(command db.CommandUnlock) db.UpdateCommandUnlockParams {
	return db.UpdateCommandUnlockParams{
		ID:     command.ID,
		Status: command.Status,
	}
}

func (r *CommandResolver) CreateCommandUnlockPayload(ctx context.Context, command db.CommandUnlock) *CommandUnlockPayload {
	return NewCommandUnlockPayload(command)
}
