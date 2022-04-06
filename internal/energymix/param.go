package energymix

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateEnergyMixParams(dto *EnergyMixDto) db.CreateEnergyMixParams {
	return db.CreateEnergyMixParams{
		IsGreenEnergy:     dto.IsGreenEnergy,
		SupplierName:      util.SqlNullString(dto.SupplierName),
		EnergyProductName: util.SqlNullString(dto.EnergyProductName),
	}
}

func NewUpdateEnergyMixParams(id int64, dto *EnergyMixDto) db.UpdateEnergyMixParams {
	return db.UpdateEnergyMixParams{
		ID:                id,
		IsGreenEnergy:     dto.IsGreenEnergy,
		SupplierName:      util.SqlNullString(dto.SupplierName),
		EnergyProductName: util.SqlNullString(dto.EnergyProductName),
	}
}

func NewCreateEnergySourceParams(id int64, dto *EnergySourceDto) db.CreateEnergySourceParams {
	return db.CreateEnergySourceParams{
		EnergyMixID: id,
		Source:      dto.Source,
		Percentage:  dto.Percentage,
	}
}

func NewCreateEnvironmentalImpactParams(id int64, dto *EnvironmentalImpactDto) db.CreateEnvironmentalImpactParams {
	return db.CreateEnvironmentalImpactParams{
		EnergyMixID: id,
		Source:      dto.Source,
		Amount:      dto.Amount,
	}
}