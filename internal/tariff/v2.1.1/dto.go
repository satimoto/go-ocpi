package tariff

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *TariffResolver) CreateTariffDto(ctx context.Context, tariff db.Tariff) *dto.TariffDto {
	response := dto.NewTariffDto(tariff)

	tariffAltTexts, err := r.Repository.ListTariffAltTexts(ctx, tariff.ID)

	if err != nil {
		util.LogOnError("OCPI256", "Error listing tariff alt texts", err)
		log.Printf("OCPI256: TariffID=%v", tariff.ID)
	} else {
		response.TariffAltText = r.DisplayTextResolver.CreateDisplayTextListDto(ctx, tariffAltTexts)
	}

	elements, err := r.ElementResolver.Repository.ListElements(ctx, tariff.ID)

	if err != nil {
		util.LogOnError("OCPI257", "Error listing elements", err)
		log.Printf("OCPI257: TariffID=%v", tariff.ID)
	} else {
		response.Elements = r.ElementResolver.CreateElementListDto(ctx, elements)
	}

	if tariff.EnergyMixID.Valid {
		energyMix, err := r.EnergyMixResolver.Repository.GetEnergyMix(ctx, tariff.EnergyMixID.Int64)

		if err != nil {
			util.LogOnError("OCPI258", "Error retrieving energy mix", err)
			log.Printf("OCPI258: EnergyMixID=%#v", tariff.EnergyMixID)
		} else {
			response.EnergyMix = r.EnergyMixResolver.CreateEnergyMixDto(ctx, energyMix)
		}
	}

	if tariff.TariffRestrictionID.Valid {
		tariffRestriction, err := r.TariffRestrictionResolver.Repository.GetTariffRestriction(ctx, tariff.TariffRestrictionID.Int64)

		if err != nil {
			util.LogOnError("OCPI259", "Error retrieving tariff restriction", err)
			log.Printf("OCPI259: TariffRestrictionID=%#v", tariff.TariffRestrictionID)
		} else {
			response.Restriction = r.TariffRestrictionResolver.CreateTariffRestrictionDto(ctx, tariffRestriction)
		}
	}

	return response
}

func (r *TariffResolver) CreateTariffPushListDto(ctx context.Context, tariffs []db.Tariff) []*dto.TariffDto {
	list := []*dto.TariffDto{}

	for _, tariff := range tariffs {
		list = append(list, r.CreateTariffDto(ctx, tariff))
	}

	return list
}
