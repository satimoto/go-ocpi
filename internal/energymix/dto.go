package energymix

import (
	"context"
	"log"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/ocpitype"
)

type EnergyMixDto struct {
	IsGreenEnergy     ocpitype.Bool             `json:"is_green_energy"`
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
		IsGreenEnergy:     ocpitype.NewBool(energyMix.IsGreenEnergy),
		SupplierName:      util.NilString(energyMix.SupplierName),
		EnergyProductName: util.NilString(energyMix.EnergyProductName),
	}
}

func (r *EnergyMixResolver) CreateEnergyMixDto(ctx context.Context, energyMix db.EnergyMix) *EnergyMixDto {
	response := NewEnergyMixDto(energyMix)

	energySources, err := r.Repository.ListEnergySources(ctx, energyMix.ID)

	if err != nil {
		util.LogOnError("OCPI229", "Error listing energy sources", err)
		log.Printf("OCPI229: EnergyMixID=%v", energyMix.ID)
	} else {
		response.EnergySources = r.CreateEnergySourceListDto(ctx, energySources)
	}

	environImpacts, err := r.Repository.ListEnvironmentalImpacts(ctx, energyMix.ID)

	if err != nil {
		util.LogOnError("OCPI230", "Error listing environmental impacts", err)
		log.Printf("OCPI230: EnergyMixID=%v", energyMix.ID)
	} else {
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
