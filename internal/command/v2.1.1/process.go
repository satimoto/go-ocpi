package command

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
)

var API_VERSION = "2.1.1"

func (r *CommandResolver) UpdateCommandReservation(ctx context.Context, command db.CommandReservation, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandReservationParams(command)
		commandParams.Status = *dto.Result

		_, err := r.Repository.UpdateCommandReservation(ctx, commandParams)

		if err != nil {
			util.LogOnError("OCPI038", "Error updating command reservation", err)
			log.Printf("OCPI038: Params=%#v", commandParams)
		}
	}
}

func (r *CommandResolver) UpdateCommandStart(ctx context.Context, command db.CommandStart, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandStartParams(command)
		commandParams.Status = *dto.Result

		_, err := r.Repository.UpdateCommandStart(ctx, commandParams)

		if err != nil {
			util.LogOnError("OCPI039", "Error updating command start", err)
			log.Printf("OCPI039: Params=%#v", commandParams)
		}
	}
}

func (r *CommandResolver) UpdateCommandStop(ctx context.Context, command db.CommandStop, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandStopParams(command)
		commandParams.Status = *dto.Result

		_, err := r.Repository.UpdateCommandStop(ctx, commandParams)

		if err != nil {
			util.LogOnError("OCPI040", "Error updating command stop", err)
			log.Printf("OCPI040: Params=%#v", commandParams)
		}
	}
}

func (r *CommandResolver) UpdateCommandUnlock(ctx context.Context, command db.CommandUnlock, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandUnlockParams(command)
		commandParams.Status = *dto.Result

		_, err := r.Repository.UpdateCommandUnlock(ctx, commandParams)

		if err != nil {
			util.LogOnError("OCPI041", "Error updating command unlock", err)
			log.Printf("OCPI041: Params=%#v", commandParams)
		}
	}
}
