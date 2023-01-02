package command

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/metric"
)

func (r *CommandResolver) CreateCommandReservationDto(ctx context.Context, command db.CommandReservation) *dto.CommandReservationDto {
	response := dto.NewCommandReservationDto(command)

	token, err := r.TokenResolver.Repository.GetToken(ctx, command.TokenID)

	if err != nil {
		metrics.RecordError("OCPI224", "Error retrieving token", err)
		log.Printf("OCPI224: TokenID=%v", command.TokenID)
		return response
	}

	response.Token = r.TokenResolver.CreateTokenDto(ctx, token)

	return response
}

func (r *CommandResolver) CreateCommandStartDto(ctx context.Context, command db.CommandStart) *dto.CommandStartDto {
	response := dto.NewCommandStartDto(command)

	token, err := r.TokenResolver.Repository.GetToken(ctx, command.TokenID)

	if err != nil {
		metrics.RecordError("OCPI225", "Error retrieving token", err)
		log.Printf("OCPI225: TokenID=%v", command.TokenID)
		return response
	}

	response.Token = r.TokenResolver.CreateTokenDto(ctx, token)

	return response
}

func (r *CommandResolver) CreateCommandStopDto(ctx context.Context, command db.CommandStop) *dto.CommandStopDto {
	return dto.NewCommandStopDto(command)
}

func (r *CommandResolver) CreateCommandUnlockDto(ctx context.Context, command db.CommandUnlock) *dto.CommandUnlockDto {
	return dto.NewCommandUnlockDto(command)
}
