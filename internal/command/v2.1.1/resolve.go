package command

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

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
	*transportation.OCPIRequester
	*session.SessionResolver
	*token.TokenResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CommandResolver {
	repo := CommandRepository(repositoryService)
	return &CommandResolver{
		Repository:            repo,
		OCPIRequester:         transportation.NewOCPIRequester(),
		SessionResolver:       session.NewResolver(repositoryService),
		TokenResolver:         token.NewResolver(repositoryService),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService),
	}
}
