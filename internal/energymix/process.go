package energymix

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *EnergyMixResolver) ReplaceEnergyMix(ctx context.Context, id *sql.NullInt64, dto *EnergyMixDto) {
	if dto != nil {
		if id.Valid {
			energyMixParams := NewUpdateEnergyMixParams(id.Int64, dto)
			_, err := r.Repository.UpdateEnergyMix(ctx, energyMixParams)

			if err != nil {
				util.LogOnError("OCPI096", "Error updating energy mix", err)
				log.Printf("OCPI096: Params=%#v", energyMixParams)
			}
		} else {
			energyMixParams := NewCreateEnergyMixParams(dto)
			energyMix, err := r.Repository.CreateEnergyMix(ctx, energyMixParams)

			if err != nil {
				util.LogOnError("OCPI095", "Error creating energy mix", err)
				log.Printf("OCPI095: Params=%#v", energyMixParams)
				return
			}

			id.Scan(energyMix.ID)
		}

		r.ReplaceEnergySources(ctx, id.Int64, *dto)
		r.ReplaceEnvironmentalImpacts(ctx, id.Int64, *dto)
	}
}

func (r *EnergyMixResolver) ReplaceEnergySources(ctx context.Context, energyMixID int64, dto EnergyMixDto) {
	r.Repository.DeleteEnergySources(ctx, energyMixID)

	for _, energySource := range dto.EnergySources {
		energySourceParams := NewCreateEnergySourceParams(energyMixID, energySource)
		_, err := r.Repository.CreateEnergySource(ctx, energySourceParams)

		if err != nil {
			util.LogOnError("OCPI097", "Error creating energy source", err)
			log.Printf("OCPI097: Params=%#v", energySourceParams)
		}
	}
}

func (r *EnergyMixResolver) ReplaceEnvironmentalImpacts(ctx context.Context, energyMixID int64, dto EnergyMixDto) {
	r.Repository.DeleteEnvironmentalImpacts(ctx, energyMixID)

	for _, environImpact := range dto.EnvironImpact {
		environImpactParams := NewCreateEnvironmentalImpactParams(energyMixID, environImpact)
		_, err := r.Repository.CreateEnvironmentalImpact(ctx, environImpactParams)

		if err != nil {
			util.LogOnError("OCPI098", "Error creating environmental impact", err)
			log.Printf("OCPI098: Params=%#v", environImpactParams)
		}
	}
}
