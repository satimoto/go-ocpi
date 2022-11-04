package energymix

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *EnergyMixResolver) CreateEnergyMixDto(ctx context.Context, energyMix db.EnergyMix) *coreDto.EnergyMixDto {
	response := coreDto.NewEnergyMixDto(energyMix)

	energySources, err := r.Repository.ListEnergySources(ctx, energyMix.ID)

	if err != nil {
		metrics.RecordError("OCPI229", "Error listing energy sources", err)
		log.Printf("OCPI229: EnergyMixID=%v", energyMix.ID)
	} else {
		response.EnergySources = r.CreateEnergySourceListDto(ctx, energySources)
	}

	environImpacts, err := r.Repository.ListEnvironmentalImpacts(ctx, energyMix.ID)

	if err != nil {
		metrics.RecordError("OCPI230", "Error listing environmental impacts", err)
		log.Printf("OCPI230: EnergyMixID=%v", energyMix.ID)
	} else {
		response.EnvironImpact = r.CreateEnvironmentalImpactListDto(ctx, environImpacts)
	}

	return response
}

func (r *EnergyMixResolver) CreateEnergySourceDto(ctx context.Context, energySource db.EnergySource) *coreDto.EnergySourceDto {
	return coreDto.NewEnergySourceDto(energySource)
}

func (r *EnergyMixResolver) CreateEnergySourceListDto(ctx context.Context, energySources []db.EnergySource) []*coreDto.EnergySourceDto {
	list := []*coreDto.EnergySourceDto{}

	for _, energySource := range energySources {
		list = append(list, r.CreateEnergySourceDto(ctx, energySource))
	}

	return list
}

func (r *EnergyMixResolver) CreateEnvironmentalImpactDto(ctx context.Context, environImpact db.EnvironmentalImpact) *coreDto.EnvironmentalImpactDto {
	return coreDto.NewEnvironmentalImpactDto(environImpact)
}

func (r *EnergyMixResolver) CreateEnvironmentalImpactListDto(ctx context.Context, environImpacts []db.EnvironmentalImpact) []*coreDto.EnvironmentalImpactDto {
	list := []*coreDto.EnvironmentalImpactDto{}

	for _, environImpact := range environImpacts {
		list = append(list, r.CreateEnvironmentalImpactDto(ctx, environImpact))
	}

	return list
}
