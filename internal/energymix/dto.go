package energymix

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

type EnergyMixDto struct {
	IsGreenEnergy     bool                      `json:"is_green_energy"`
	EnergySources     []*EnergySourceDto        `json:"energy_sources,omitempty"`
	EnvironImpact     []*EnvironmentalImpactDto `json:"environ_impact,omitempty"`
	SupplierName      *string                   `json:"supplier_name,omitempty"`
	EnergyProductName *string                   `json:"energy_product_name,omitempty"`
}

func (r *EnergyMixDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEnergyMixDto(energyMix db.EnergyMix) *EnergyMixDto {
	return &EnergyMixDto{
		IsGreenEnergy:     energyMix.IsGreenEnergy,
		SupplierName:      util.NilString(energyMix.SupplierName),
		EnergyProductName: util.NilString(energyMix.EnergyProductName),
	}
}

func (r *EnergyMixResolver) CreateEnergyMixDto(ctx context.Context, energyMix db.EnergyMix) *EnergyMixDto {
	response := NewEnergyMixDto(energyMix)

	if energySources, err := r.Repository.ListEnergySources(ctx, energyMix.ID); err == nil {
		response.EnergySources = r.CreateEnergySourceListDto(ctx, energySources)
	}

	if environImpacts, err := r.Repository.ListEnvironmentalImpacts(ctx, energyMix.ID); err == nil {
		response.EnvironImpact = r.CreateEnvironmentalImpactListDto(ctx, environImpacts)
	}

	return response
}

type EnergySourceDto struct {
	Source     db.EnergySourceCategory `json:"source"`
	Percentage float64                 `json:"percentage"`
}

func (r *EnergySourceDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEnergySourceDto(energySource db.EnergySource) *EnergySourceDto {
	return &EnergySourceDto{
		Source:     energySource.Source,
		Percentage: energySource.Percentage,
	}
}

func (r *EnergyMixResolver) CreateEnergySourceDto(ctx context.Context, energySource db.EnergySource) *EnergySourceDto {
	return NewEnergySourceDto(energySource)
}

func (r *EnergyMixResolver) CreateEnergySourceListDto(ctx context.Context, energySources []db.EnergySource) []*EnergySourceDto {
	list := []*EnergySourceDto{}
	for _, energySource := range energySources {
		list = append(list, r.CreateEnergySourceDto(ctx, energySource))
	}
	return list
}

type EnvironmentalImpactDto struct {
	Source db.EnvironmentalImpactCategory `json:"source"`
	Amount float64                        `json:"amount"`
}

func (r *EnvironmentalImpactDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewEnvironmentalImpactDto(environmentalImpact db.EnvironmentalImpact) *EnvironmentalImpactDto {
	return &EnvironmentalImpactDto{
		Source: environmentalImpact.Source,
		Amount: environmentalImpact.Amount,
	}
}

func (r *EnergyMixResolver) CreateEnvironmentalImpactDto(ctx context.Context, environImpact db.EnvironmentalImpact) *EnvironmentalImpactDto {
	return NewEnvironmentalImpactDto(environImpact)
}

func (r *EnergyMixResolver) CreateEnvironmentalImpactListDto(ctx context.Context, environImpacts []db.EnvironmentalImpact) []*EnvironmentalImpactDto {
	list := []*EnvironmentalImpactDto{}
	for _, environImpact := range environImpacts {
		list = append(list, r.CreateEnvironmentalImpactDto(ctx, environImpact))
	}
	return list
}
