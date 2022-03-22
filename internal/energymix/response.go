package energymix

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type EnergyMixPayload struct {
	ID                int64   `json:"id"`
	IsGreenEnergy     bool    `json:"is_green_energy"`
	SupplierName      *string `json:"supplier_name,omitempty"`
	EnergyProductName *string `json:"energy_product_name,omitempty"`
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
	return NewEnergyMixPayload(energyMix)
}
