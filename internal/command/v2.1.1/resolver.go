package command

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
)

var API_VERSION = "2.1.1"

type CommandRepository interface {
	CreateCommandReservation(ctx context.Context, arg db.CreateCommandReservationParams) (db.CommandReservation, error)
	CreateCommandStart(ctx context.Context, arg db.CreateCommandStartParams) (db.CommandStart, error)
	CreateCommandStop(ctx context.Context, arg db.CreateCommandStopParams) (db.CommandStop, error)
	CreateCommandUnlock(ctx context.Context, arg db.CreateCommandUnlockParams) (db.CommandUnlock, error)
	GetCommandReservation(ctx context.Context, id int64) (db.CommandReservation, error)
	GetCommandStart(ctx context.Context, id int64) (db.CommandStart, error)
	GetCommandStop(ctx context.Context, id int64) (db.CommandStop, error)
	GetCommandUnlock(ctx context.Context, id int64) (db.CommandUnlock, error)
	UpdateCommandReservation(ctx context.Context, arg db.UpdateCommandReservationParams) (db.CommandReservation, error)
	UpdateCommandStart(ctx context.Context, arg db.UpdateCommandStartParams) (db.CommandStart, error)
	UpdateCommandStop(ctx context.Context, arg db.UpdateCommandStopParams) (db.CommandStop, error)
	UpdateCommandUnlock(ctx context.Context, arg db.UpdateCommandUnlockParams) (db.CommandUnlock, error)
}

type CommandResolver struct {
	Repository CommandRepository
	*token.TokenResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CommandResolver {
	repo := CommandRepository(repositoryService)
	return &CommandResolver{
		Repository:    repo,
		TokenResolver: token.NewResolver(repositoryService),
	}
}

func (r *CommandResolver) UpdateCommandReservation(ctx context.Context, command db.CommandReservation, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := NewUpdateCommandReservationParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandReservation(ctx, commandParams)
	}
}

func (r *CommandResolver) UpdateCommandStart(ctx context.Context, command db.CommandStart, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := NewUpdateCommandStartParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandStart(ctx, commandParams)
	}
}

func (r *CommandResolver) UpdateCommandStop(ctx context.Context, command db.CommandStop, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := NewUpdateCommandStopParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandStop(ctx, commandParams)
	}
}

func (r *CommandResolver) UpdateCommandUnlock(ctx context.Context, command db.CommandUnlock, dto *CommandResponseDto) {
	if dto != nil {
		commandParams := NewUpdateCommandUnlockParams(command)
		commandParams.Status = *dto.Result

		r.Repository.UpdateCommandUnlock(ctx, commandParams)
	}
}
