package openingtime

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *OpeningTimeResolver) CreateExceptionalPeriodDto(ctx context.Context, exceptionalPeriod db.ExceptionalPeriod) *coreDto.ExceptionalPeriodDto {
	return coreDto.NewExceptionalPeriodDto(exceptionalPeriod)
}

func (r *OpeningTimeResolver) CreateExceptionalPeriodListDto(ctx context.Context, exceptionalPeriods []db.ExceptionalPeriod) []*coreDto.ExceptionalPeriodDto {
	list := []*coreDto.ExceptionalPeriodDto{}

	for _, exceptionalPeriod := range exceptionalPeriods {
		list = append(list, r.CreateExceptionalPeriodDto(ctx, exceptionalPeriod))
	}

	return list
}

func (r *OpeningTimeResolver) CreateOpeningTimeDto(ctx context.Context, openingTime db.OpeningTime) *coreDto.OpeningTimeDto {
	response := coreDto.NewOpeningTimeDto(openingTime)

	regularHours, err := r.Repository.ListRegularHours(ctx, openingTime.ID)

	if err != nil {
		metrics.RecordError("OCPI249", "Error listing regular hours", err)
		log.Printf("OCPI249: OpeningTimeID=%v", openingTime.ID)
	} else {
		response.RegularHours = r.CreateRegularHourListDto(ctx, regularHours)
	}

	exceptionalOpenings, err := r.Repository.ListExceptionalOpeningPeriods(ctx, openingTime.ID)

	if err != nil {
		metrics.RecordError("OCPI250", "Error listing exceptional opening periods", err)
		log.Printf("OCPI250: OpeningTimeID=%v", openingTime.ID)
	} else {
		response.ExceptionalOpenings = r.CreateExceptionalPeriodListDto(ctx, exceptionalOpenings)
	}

	exceptionalClosings, err := r.Repository.ListExceptionalClosingPeriods(ctx, openingTime.ID)

	if err != nil {
		metrics.RecordError("OCPI251", "Error listing exceptional closing periods", err)
		log.Printf("OCPI251: OpeningTimeID=%v", openingTime.ID)
	} else {
		response.ExceptionalClosings = r.CreateExceptionalPeriodListDto(ctx, exceptionalClosings)
	}

	return response
}

func (r *OpeningTimeResolver) CreateRegularHourDto(ctx context.Context, regularHour db.RegularHour) *coreDto.RegularHourDto {
	return coreDto.NewRegularHourDto(regularHour)
}

func (r *OpeningTimeResolver) CreateRegularHourListDto(ctx context.Context, regularHours []db.RegularHour) []*coreDto.RegularHourDto {
	list := []*coreDto.RegularHourDto{}

	for _, regularHour := range regularHours {
		list = append(list, r.CreateRegularHourDto(ctx, regularHour))
	}

	return list
}
