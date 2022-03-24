package energymix

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type EnergyMixPayload struct {
	IsGreenEnergy     bool                          `json:"is_green_energy"`
	EnergySources     []*EnergySourcePayload        `json:"energy_sources,omitempty"`
	EnvironImpact     []*EnvironmentalImpactPayload `json:"environ_impact,omitempty"`
	SupplierName      *string                       `json:"supplier_name,omitempty"`
	EnergyProductName *string                       `json:"energy_product_name,omitempty"`
}

func (r *EnergyMixPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEnergyMixPayload(energyMix db.EnergyMix) *EnergyMixPayload {
	return &EnergyMixPayload{
		IsGreenEnergy:     energyMix.IsGreenEnergy,
		SupplierName:      util.NilString(energyMix.SupplierName.String),
		EnergyProductName: util.NilString(energyMix.EnergyProductName.String),
	}
}

func NewCreateEnergyMixParams(payload *EnergyMixPayload) db.CreateEnergyMixParams {
	return db.CreateEnergyMixParams{
		IsGreenEnergy:     payload.IsGreenEnergy,
		SupplierName:      util.SqlNullString(payload.SupplierName),
		EnergyProductName: util.SqlNullString(payload.EnergyProductName),
	}
}

func NewUpdateEnergyMixParams(id int64, payload *EnergyMixPayload) db.UpdateEnergyMixParams {
	return db.UpdateEnergyMixParams{
		ID:                id,
		IsGreenEnergy:     payload.IsGreenEnergy,
		SupplierName:      util.SqlNullString(payload.SupplierName),
		EnergyProductName: util.SqlNullString(payload.EnergyProductName),
	}
}

func (r *EnergyMixResolver) CreateEnergyMixPayload(ctx context.Context, energyMix db.EnergyMix) *EnergyMixPayload {
	response := NewEnergyMixPayload(energyMix)

	if energySources, err := r.Repository.ListEnergySources(ctx, energyMix.ID); err == nil {
		response.EnergySources = r.CreateEnergySourceListPayload(ctx, energySources)
	}

	if environImpacts, err := r.Repository.ListEnvironmentalImpacts(ctx, energyMix.ID); err == nil {
		response.EnvironImpact = r.CreateEnvironmentalImpactListPayload(ctx, environImpacts)
	}

	return response
}

type EnergySourcePayload struct {
	Source     db.EnergySourceCategory `json:"source"`
	Percentage float64                 `json:"percentage"`
}

func (r *EnergySourcePayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEnergySourcePayload(energySource db.EnergySource) *EnergySourcePayload {
	return &EnergySourcePayload{
		Source:     energySource.Source,
		Percentage: energySource.Percentage,
	}
}

func NewCreateEnergySourceParams(id int64, payload *EnergySourcePayload) db.CreateEnergySourceParams {
	return db.CreateEnergySourceParams{
		EnergyMixID: id,
		Source:      payload.Source,
		Percentage:  payload.Percentage,
	}
}

func (r *EnergyMixResolver) CreateEnergySourcePayload(ctx context.Context, energySource db.EnergySource) *EnergySourcePayload {
	return NewEnergySourcePayload(energySource)
}

func (r *EnergyMixResolver) CreateEnergySourceListPayload(ctx context.Context, energySources []db.EnergySource) []*EnergySourcePayload {
	list := []*EnergySourcePayload{}
	for _, energySource := range energySources {
		list = append(list, r.CreateEnergySourcePayload(ctx, energySource))
	}
	return list
}

type EnvironmentalImpactPayload struct {
	Source db.EnvironmentalImpactCategory `json:"source"`
	Amount float64                        `json:"amount"`
}

func (r *EnvironmentalImpactPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEnvironmentalImpactPayload(environmentalImpact db.EnvironmentalImpact) *EnvironmentalImpactPayload {
	return &EnvironmentalImpactPayload{
		Source: environmentalImpact.Source,
		Amount: environmentalImpact.Amount,
	}
}

func NewCreateEnvironmentalImpactParams(id int64, payload *EnvironmentalImpactPayload) db.CreateEnvironmentalImpactParams {
	return db.CreateEnvironmentalImpactParams{
		EnergyMixID: id,
		Source:      payload.Source,
		Amount:      payload.Amount,
	}
}

func (r *EnergyMixResolver) CreateEnvironmentalImpactPayload(ctx context.Context, environImpact db.EnvironmentalImpact) *EnvironmentalImpactPayload {
	return NewEnvironmentalImpactPayload(environImpact)
}

func (r *EnergyMixResolver) CreateEnvironmentalImpactListPayload(ctx context.Context, environImpacts []db.EnvironmentalImpact) []*EnvironmentalImpactPayload {
	list := []*EnvironmentalImpactPayload{}
	for _, environImpact := range environImpacts {
		list = append(list, r.CreateEnvironmentalImpactPayload(ctx, environImpact))
	}
	return list
}
