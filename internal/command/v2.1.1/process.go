package command

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *CommandResolver) UpdateCommandReservation(ctx context.Context, command db.CommandReservation, commandResponseDto *dto.CommandResponseDto) {
	if commandResponseDto != nil {
		commandParams := param.NewUpdateCommandReservationParams(command)
		commandParams.Status = *commandResponseDto.Result
		commandParams.LastUpdated = time.Now().UTC()

		_, err := r.Repository.UpdateCommandReservation(ctx, commandParams)

		if err != nil {
			metrics.RecordError("OCPI038", "Error updating command reservation", err)
			log.Printf("OCPI038: Params=%#v", commandParams)
		}
	}
}

func (r *CommandResolver) UpdateCommandStart(ctx context.Context, credential db.Credential, command db.CommandStart, commandResponseDto *dto.CommandResponseDto) {
	if commandResponseDto != nil {
		commandParams := param.NewUpdateCommandStartParams(command)
		commandParams.Status = *commandResponseDto.Result
		commandParams.LastUpdated = time.Now().UTC()

		updatedCommand, err := r.Repository.UpdateCommandStart(ctx, commandParams)

		if err != nil {
			metrics.RecordError("OCPI039", "Error updating command start", err)
			log.Printf("OCPI039: Params=%#v", commandParams)
		}

		if updatedCommand.Status == db.CommandResponseTypeACCEPTED && updatedCommand.AuthorizationID.Valid {
			go r.waitForEvseStatus(credential, updatedCommand)
		}
	}
}

func (r *CommandResolver) UpdateCommandStop(ctx context.Context, command db.CommandStop, commandResponseDto *dto.CommandResponseDto) {
	if commandResponseDto != nil {
		commandParams := param.NewUpdateCommandStopParams(command)
		commandParams.Status = *commandResponseDto.Result
		commandParams.LastUpdated = time.Now().UTC()

		_, err := r.Repository.UpdateCommandStop(ctx, commandParams)

		if err != nil {
			metrics.RecordError("OCPI040", "Error updating command stop", err)
			log.Printf("OCPI040: Params=%#v", commandParams)
		}
	}
}

func (r *CommandResolver) UpdateCommandUnlock(ctx context.Context, command db.CommandUnlock, commandResponseDto *dto.CommandResponseDto) {
	if commandResponseDto != nil {
		commandParams := param.NewUpdateCommandUnlockParams(command)
		commandParams.Status = *commandResponseDto.Result
		commandParams.LastUpdated = time.Now().UTC()

		_, err := r.Repository.UpdateCommandUnlock(ctx, commandParams)

		if err != nil {
			metrics.RecordError("OCPI041", "Error updating command unlock", err)
			log.Printf("OCPI041: Params=%#v", commandParams)
		}
	}
}

func (r *CommandResolver) waitForEvseStatus(credential db.Credential, command db.CommandStart) {
	ctx := context.Background()
	token, err := r.TokenResolver.Repository.GetToken(ctx, command.TokenID)

	if err != nil {
		metrics.RecordError("OCPI325", "Error getting token", err)
		log.Printf("OCPI325: CommandID=%v, TokenID=%v", command.ID, command.TokenID)
		return
	}

	tokenAuthorization, err := r.TokenResolver.TokenAuthorizationResolver.Repository.GetTokenAuthorizationByAuthorizationID(ctx, command.AuthorizationID.String)

	if err != nil {
		metrics.RecordError("OCPI326", "Error updating command start", err)
		log.Printf("OCPI326: CommandID=%v, AuthorizationID=%#v", command.ID, command.AuthorizationID)
		return
	}

	if command.EvseUid.Valid {
		cancelCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		r.EvseResolver.WaitForEvseStatus(credential, token, tokenAuthorization, command.LocationID, command.EvseUid.String, db.EvseStatusCHARGING, cancelCtx, cancel, 150)

		<-cancelCtx.Done()
	}
}
