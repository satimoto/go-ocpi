package session

import (
	"context"
	"log"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *SessionResolver) CreateSessionDto(ctx context.Context, session db.Session) *dto.SessionDto {
	response := dto.NewSessionDto(session)

	chargingPeriods, err := r.Repository.ListSessionChargingPeriods(ctx, session.ID)

	if err != nil {
		metrics.RecordError("OCPI254", "Error listing session charging periods", err)
		log.Printf("OCPI254: SessionID=%v", session.ID)
	} else {
		response.ChargingPeriods = r.ChargingPeriodResolver.CreateChargingPeriodListDto(ctx, chargingPeriods)
	}

	location, err := r.LocationResolver.Repository.GetLocation(ctx, session.LocationID)

	if err != nil {
		metrics.RecordError("OCPI255", "Error listing session charging periods", err)
		log.Printf("OCPI255: LocationID=%v", session.LocationID)
	} else {
		response.Location = r.LocationResolver.CreateLocationDto(ctx, location)
	}

	return response
}

func (r *SessionResolver) CreateSessionListDto(ctx context.Context, sessions []db.Session) []render.Renderer {
	list := []render.Renderer{}

	for _, session := range sessions {
		list = append(list, r.CreateSessionDto(ctx, session))
	}

	return list
}
