package energymix

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type EnergyMixRepository interface {
	CreateEnergyMix(ctx context.Context, arg db.CreateEnergyMixParams) (db.EnergyMix, error)
	CreateEnergySource(ctx context.Context, arg db.CreateEnergySourceParams) (db.EnergySource, error)
	CreateEnvironmentalImpact(ctx context.Context, arg db.CreateEnvironmentalImpactParams) (db.EnvironmentalImpact, error)
	DeleteEnergySources(ctx context.Context, energyMixID int64) error
	DeleteEnvironmentalImpacts(ctx context.Context, energyMixID int64) error
	GetEnergyMix(ctx context.Context, id int64) (db.EnergyMix, error)
	ListEnergySources(ctx context.Context, energyMixID int64) ([]db.EnergySource, error)
	ListEnvironmentalImpacts(ctx context.Context, energyMixID int64) ([]db.EnvironmentalImpact, error)
	UpdateEnergyMix(ctx context.Context, arg db.UpdateEnergyMixParams) (db.EnergyMix, error)
}

type EnergyMixResolver struct {
	Repository EnergyMixRepository
}

func NewResolver(repositoryService *db.RepositoryService) *EnergyMixResolver {
	repo := EnergyMixRepository(repositoryService)
	
	return &EnergyMixResolver{repo}
}
