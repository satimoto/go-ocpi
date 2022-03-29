package chargingperiod

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
)

type ChargingPeriodPayload struct {
	StartDateTime *time.Time                        `json:"start_date_time"`
	Dimensions    []*ChargingPeriodDimensionPayload `json:"dimensions"`
}

func (r *ChargingPeriodPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewChargingPeriodPayload(chargingPeriod db.ChargingPeriod) *ChargingPeriodPayload {
	return &ChargingPeriodPayload{
		StartDateTime: &chargingPeriod.StartDateTime,
	}
}

func (r *ChargingPeriodResolver) CreateChargingPeriodPayload(ctx context.Context, chargingPeriod db.ChargingPeriod) *ChargingPeriodPayload {
	response := NewChargingPeriodPayload(chargingPeriod)

	if chargingPeriodDimensions, err := r.Repository.ListChargingPeriodDimensions(ctx, chargingPeriod.ID); err == nil {
		response.Dimensions = r.CreateChargingPeriodDimensionListPayload(ctx, chargingPeriodDimensions)
	}

	return response
}

func (r *ChargingPeriodResolver) CreateChargingPeriodListPayload(ctx context.Context, chargingPeriods []db.ChargingPeriod) []*ChargingPeriodPayload {
	list := []*ChargingPeriodPayload{}
	for _, chargingPeriod := range chargingPeriods {
		list = append(list, r.CreateChargingPeriodPayload(ctx, chargingPeriod))
	}
	return list
}

type ChargingPeriodDimensionPayload struct {
	Type   db.ChargingPeriodDimensionType `json:"type"`
	Volume float64                        `json:"volume"`
}

func (r *ChargingPeriodDimensionPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewChargingPeriodDimensionPayload(chargingPeriodDimension db.ChargingPeriodDimension) *ChargingPeriodDimensionPayload {
	return &ChargingPeriodDimensionPayload{
		Type:   chargingPeriodDimension.Type,
		Volume: chargingPeriodDimension.Volume,
	}
}

func NewCreateChargingPeriodDimensionParams(id int64, payload *ChargingPeriodDimensionPayload) db.CreateChargingPeriodDimensionParams {
	return db.CreateChargingPeriodDimensionParams{
		ChargingPeriodID: id,
		Type:             payload.Type,
		Volume:           payload.Volume,
	}
}

func (r *ChargingPeriodResolver) CreateChargingPeriodDimensionPayload(ctx context.Context, chargingPeriodDimension db.ChargingPeriodDimension) *ChargingPeriodDimensionPayload {
	return NewChargingPeriodDimensionPayload(chargingPeriodDimension)
}

func (r *ChargingPeriodResolver) CreateChargingPeriodDimensionListPayload(ctx context.Context, chargingPeriodDimensions []db.ChargingPeriodDimension) []*ChargingPeriodDimensionPayload {
	list := []*ChargingPeriodDimensionPayload{}
	for _, chargingPeriodDimension := range chargingPeriodDimensions {
		list = append(list, r.CreateChargingPeriodDimensionPayload(ctx, chargingPeriodDimension))
	}
	return list
}
