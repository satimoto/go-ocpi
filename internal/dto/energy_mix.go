package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
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
