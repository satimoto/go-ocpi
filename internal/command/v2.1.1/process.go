package command

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
)

var API_VERSION = "2.1.1"

func (r *CommandResolver) UpdateCommandReservation(ctx context.Context, command db.CommandReservation, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandReservationParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandReservation(ctx, commandParams)
	}
}

func (r *CommandResolver) UpdateCommandStart(ctx context.Context, command db.CommandStart, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandStartParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandStart(ctx, commandParams)
	}
}

func (r *CommandResolver) UpdateCommandStop(ctx context.Context, command db.CommandStop, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandStopParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandStop(ctx, commandParams)
	}
}

func (r *CommandResolver) UpdateCommandUnlock(ctx context.Context, command db.CommandUnlock, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := param.NewUpdateCommandUnlockParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandUnlock(ctx, commandParams)
	}
}
