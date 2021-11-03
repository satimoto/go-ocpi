package energymix

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EnergyMixRepository interface {
	CreateEnergyMix(ctx context.Context, arg db.CreateEnergyMixParams) (db.EnergyMix, error)
	GetEnergyMix(ctx context.Context, id int64) (db.EnergyMix, error)
	UpdateEnergyMix(ctx context.Context, arg db.UpdateEnergyMixParams) (db.EnergyMix, error)
}

type EnergyMixResolver struct {
	Repository EnergyMixRepository
}

func NewResolver(repositoryService *db.RepositoryService) *EnergyMixResolver {
	repo := EnergyMixRepository(repositoryService)
	return &EnergyMixResolver{repo}
}

func (r *EnergyMixResolver) ReplaceEnergyMix(ctx context.Context, id *int64, payload *EnergyMixPayload) {
	if payload != nil {
		if id == nil {
			energyMixParams := NewCreateEnergyMixParams(payload)

			if energyMix, err := r.Repository.CreateEnergyMix(ctx, energyMixParams); err == nil {
				id = &energyMix.ID
			}
		} else {
			energyMixParams := NewUpdateEnergyMixParams(*id, payload)

			r.Repository.UpdateEnergyMix(ctx, energyMixParams)	
		}
	}
}
