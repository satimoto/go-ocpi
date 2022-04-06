package energymix

import (
	"context"
)

func (r *EnergyMixResolver) ReplaceEnergyMix(ctx context.Context, id *int64, dto *EnergyMixDto) {
	if dto != nil {
		if id == nil {
			energyMixParams := NewCreateEnergyMixParams(dto)

			if energyMix, err := r.Repository.CreateEnergyMix(ctx, energyMixParams); err == nil {
				id = &energyMix.ID
			}
		} else {
			energyMixParams := NewUpdateEnergyMixParams(*id, dto)

			r.Repository.UpdateEnergyMix(ctx, energyMixParams)
		}

		r.ReplaceEnergySources(ctx, id, *dto)
		r.ReplaceEnvironmentalImpacts(ctx, id, *dto)
	}
}

func (r *EnergyMixResolver) ReplaceEnergySources(ctx context.Context, energyMixID *int64, dto EnergyMixDto) {
	if energyMixID != nil {
		r.Repository.DeleteEnergySources(ctx, *energyMixID)

		for _, energySource := range dto.EnergySources {
			energySourceParams := NewCreateEnergySourceParams(*energyMixID, energySource)
			r.Repository.CreateEnergySource(ctx, energySourceParams)
		}
	}
}

func (r *EnergyMixResolver) ReplaceEnvironmentalImpacts(ctx context.Context, energyMixID *int64, dto EnergyMixDto) {
	if energyMixID != nil {
		r.Repository.DeleteEnvironmentalImpacts(ctx, *energyMixID)

		for _, environImpact := range dto.EnvironImpact {
			environImpactParams := NewCreateEnvironmentalImpactParams(*energyMixID, environImpact)
			r.Repository.CreateEnvironmentalImpact(ctx, environImpactParams)
		}
	}
}
