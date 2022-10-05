package energymix

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateEnergyMixParams(energyMixDto *coreDto.EnergyMixDto) db.CreateEnergyMixParams {
	return db.CreateEnergyMixParams{
		IsGreenEnergy:     energyMixDto.IsGreenEnergy.Bool(),
		SupplierName:      util.SqlNullString(energyMixDto.SupplierName),
		EnergyProductName: util.SqlNullString(energyMixDto.EnergyProductName),
	}
}

func NewUpdateEnergyMixParams(id int64, energyMixDto *coreDto.EnergyMixDto) db.UpdateEnergyMixParams {
	return db.UpdateEnergyMixParams{
		ID:                id,
		IsGreenEnergy:     energyMixDto.IsGreenEnergy.Bool(),
		SupplierName:      util.SqlNullString(energyMixDto.SupplierName),
		EnergyProductName: util.SqlNullString(energyMixDto.EnergyProductName),
	}
}

func NewCreateEnergySourceParams(id int64, energySourceDto *coreDto.EnergySourceDto) db.CreateEnergySourceParams {
	return db.CreateEnergySourceParams{
		EnergyMixID: id,
		Source:      energySourceDto.Source,
		Percentage:  energySourceDto.Percentage,
	}
}

func NewCreateEnvironmentalImpactParams(id int64, environmentalImpactDto *coreDto.EnvironmentalImpactDto) db.CreateEnvironmentalImpactParams {
	return db.CreateEnvironmentalImpactParams{
		EnergyMixID: id,
		Source:      environmentalImpactDto.Source,
		Amount:      environmentalImpactDto.Amount,
	}
}
