package openingtime

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type ExceptionalPeriodDto struct {
	PeriodBegin *time.Time `json:"period_begin"`
	PeriodEnd   *time.Time `json:"period_end"`
}

func (r *ExceptionalPeriodDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewExceptionalPeriodDto(exceptionalPeriod db.ExceptionalPeriod) *ExceptionalPeriodDto {
	return &ExceptionalPeriodDto{
		PeriodBegin: &exceptionalPeriod.PeriodBegin,
		PeriodEnd:   &exceptionalPeriod.PeriodEnd,
	}
}

func (r *OpeningTimeResolver) CreateExceptionalPeriodDto(ctx context.Context, exceptionalPeriod db.ExceptionalPeriod) *ExceptionalPeriodDto {
	return NewExceptionalPeriodDto(exceptionalPeriod)
}

func (r *OpeningTimeResolver) CreateExceptionalPeriodListDto(ctx context.Context, exceptionalPeriods []db.ExceptionalPeriod) []*ExceptionalPeriodDto {
	list := []*ExceptionalPeriodDto{}

	for _, exceptionalPeriod := range exceptionalPeriods {
		list = append(list, r.CreateExceptionalPeriodDto(ctx, exceptionalPeriod))
	}

	return list
}

type OpeningTimeDto struct {
	RegularHours        []*RegularHourDto       `json:"regular_hours"`
	Twentyfourseven     bool                    `json:"twentyfourseven"`
	ExceptionalOpenings []*ExceptionalPeriodDto `json:"exceptional_openings"`
	ExceptionalClosings []*ExceptionalPeriodDto `json:"exceptional_closings"`
}

func (r *OpeningTimeDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewOpeningTimeDto(openingTime db.OpeningTime) *OpeningTimeDto {
	return &OpeningTimeDto{
		Twentyfourseven: openingTime.Twentyfourseven,
	}
}

func (r *OpeningTimeResolver) CreateOpeningTimeDto(ctx context.Context, openingTime db.OpeningTime) *OpeningTimeDto {
	response := NewOpeningTimeDto(openingTime)

	regularHours, err := r.Repository.ListRegularHours(ctx, openingTime.ID)

	if err != nil {
		util.LogOnError("OCPI249", "Error listing regular hours", err)
		log.Printf("OCPI249: OpeningTimeID=%v", openingTime.ID)
	} else {
		response.RegularHours = r.CreateRegularHourListDto(ctx, regularHours)
	}

	exceptionalOpenings, err := r.Repository.ListExceptionalOpeningPeriods(ctx, openingTime.ID)
	
	if err != nil {
		util.LogOnError("OCPI250", "Error listing exceptional opening periods", err)
		log.Printf("OCPI250: OpeningTimeID=%v", openingTime.ID)
	} else {
		response.ExceptionalOpenings = r.CreateExceptionalPeriodListDto(ctx, exceptionalOpenings)
	}

	exceptionalClosings, err := r.Repository.ListExceptionalClosingPeriods(ctx, openingTime.ID)

	if err != nil {
		util.LogOnError("OCPI251", "Error listing exceptional closing periods", err)
		log.Printf("OCPI251: OpeningTimeID=%v", openingTime.ID)
	} else {
		response.ExceptionalClosings = r.CreateExceptionalPeriodListDto(ctx, exceptionalClosings)
	}

	return response
}

type RegularHourDto struct {
	Weekday     int16  `json:"weekday"`
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

func (r *RegularHourDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewRegularHourDto(regularHour db.RegularHour) *RegularHourDto {
	return &RegularHourDto{
		Weekday:     regularHour.Weekday,
		PeriodBegin: regularHour.PeriodBegin,
		PeriodEnd:   regularHour.PeriodEnd,
	}
}

func (r *OpeningTimeResolver) CreateRegularHourDto(ctx context.Context, regularHour db.RegularHour) *RegularHourDto {
	return NewRegularHourDto(regularHour)
}

func (r *OpeningTimeResolver) CreateRegularHourListDto(ctx context.Context, regularHours []db.RegularHour) []*RegularHourDto {
	list := []*RegularHourDto{}
	
	for _, regularHour := range regularHours {
		list = append(list, r.CreateRegularHourDto(ctx, regularHour))
	}

	return list
}
