package dto

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
	"github.com/satimoto/go-ocpi/internal/version"
)

type OcpiCommandResponseDto struct {
	Data          *CommandResponseDto `json:"data,omitempty"`
	StatusCode    int16               `json:"status_code"`
	StatusMessage string              `json:"status_message"`
	Timestamp     ocpitype.Time       `json:"timestamp"`
}

type CommandReservationDto struct {
	ResponseUrl   *string    `json:"response_url"`
	Token         *TokenDto  `json:"token"`
	ExpiryDate    *time.Time `json:"expiry_date"`
	ReservationID *int64     `json:"reservation_id"`
	LocationID    *string    `json:"location_id"`
	EvseUid       *string    `json:"evse_uid,omitempty"`
}

func (r *CommandReservationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandReservationDto(command db.CommandReservation) *CommandReservationDto {
	responseUrl := fmt.Sprintf("%s/%s/commands/RESERVE_NOW/%v", os.Getenv("API_DOMAIN"), version.VERSION_2_1_1, command.ID)

	return &CommandReservationDto{
		ResponseUrl:   &responseUrl,
		ExpiryDate:    &command.ExpiryDate,
		ReservationID: &command.ReservationID,
		LocationID:    &command.LocationID,
		EvseUid:       util.NilString(command.EvseUid),
	}
}

type CommandResponseDto struct {
	Result *db.CommandResponseType `json:"result"`
}

func (r *CommandResponseDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type CommandStartDto struct {
	ResponseUrl     *string   `json:"response_url"`
	Token           *TokenDto `json:"token"`
	AuthorizationID *string   `json:"authorization_id,omitempty"`
	LocationID      *string   `json:"location_id"`
	EvseUid         *string   `json:"evse_uid,omitempty"`
}

func (r *CommandStartDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandStartDto(command db.CommandStart) *CommandStartDto {
	responseUrl := fmt.Sprintf("%s/%s/commands/START_SESSION/%v", os.Getenv("API_DOMAIN"), version.VERSION_2_1_1, command.ID)

	return &CommandStartDto{
		ResponseUrl:     &responseUrl,
		AuthorizationID: util.NilString(command.AuthorizationID),
		LocationID:      &command.LocationID,
		EvseUid:         util.NilString(command.EvseUid),
	}
}

type CommandStopDto struct {
	ResponseUrl *string `json:"response_url"`
	SessionID   *string `json:"session_id"`
}

func (r *CommandStopDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandStopDto(command db.CommandStop) *CommandStopDto {
	responseUrl := fmt.Sprintf("%s/%s/commands/STOP_SESSION/%v", os.Getenv("API_DOMAIN"), version.VERSION_2_1_1, command.ID)

	return &CommandStopDto{
		ResponseUrl: &responseUrl,
		SessionID:   &command.SessionID,
	}
}

type CommandUnlockDto struct {
	ResponseUrl *string `json:"response_url"`
	LocationID  *string `json:"location_id"`
	EvseUid     *string `json:"evse_uid"`
	ConnectorID *string `json:"connector_id"`
}

func (r *CommandUnlockDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCommandUnlockDto(command db.CommandUnlock) *CommandUnlockDto {
	responseUrl := fmt.Sprintf("%s/%s/commands/UNLOCK_CONNECTOR/%v", os.Getenv("API_DOMAIN"), version.VERSION_2_1_1, command.ID)

	return &CommandUnlockDto{
		ResponseUrl: &responseUrl,
		LocationID:  &command.LocationID,
		EvseUid:     &command.EvseUid,
		ConnectorID: &command.ConnectorID,
	}
}
