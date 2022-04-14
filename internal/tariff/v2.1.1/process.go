package tariff

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
)

func (r *TariffResolver) ReplaceTariffByIdentifier(ctx context.Context, countryCode *string, partyID *string, uid string, cdrID *int64, dto *TariffDto) *db.Tariff {
	if dto != nil {
		tariff, err := r.Repository.GetTariffByUid(ctx, uid)
		energyMixID := util.NilInt64(tariff.EnergyMixID)
		tariffRestrictionID := util.NilInt64(tariff.TariffRestrictionID)

		if dto.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, energyMixID, dto.EnergyMix)
		}

		if dto.Restriction != nil {
			r.TariffRestrictionResolver.ReplaceTariffByIdentifierRestriction(ctx, tariffRestrictionID, dto.Restriction)
		}

		if err == nil {
			tariffParams := NewUpdateTariffByUidParams(tariff)
			tariffParams.CountryCode = util.SqlNullString(countryCode)
			tariffParams.PartyID = util.SqlNullString(partyID)
			tariffParams.EnergyMixID = util.SqlNullInt64(energyMixID)
			tariffParams.TariffRestrictionID = util.SqlNullInt64(tariffRestrictionID)

			if dto.Currency != nil {
				tariffParams.Currency = *dto.Currency
			}

			if dto.LastUpdated != nil {
				tariffParams.LastUpdated = *dto.LastUpdated
			}

			if dto.TariffAltUrl != nil {
				tariffParams.TariffAltUrl = util.SqlNullString(dto.TariffAltUrl)
			}

			tariff, err = r.Repository.UpdateTariffByUid(ctx, tariffParams)
		} else {
			tariffParams := NewCreateTariffParams(dto)
			tariffParams.CountryCode = util.SqlNullString(countryCode)
			tariffParams.PartyID = util.SqlNullString(partyID)
			tariffParams.CdrID = util.SqlNullInt64(cdrID)
			tariffParams.EnergyMixID = util.SqlNullInt64(energyMixID)
			tariffParams.TariffRestrictionID = util.SqlNullInt64(tariffRestrictionID)

			tariff, err = r.Repository.CreateTariff(ctx, tariffParams)
		}

		if dto.TariffAltText != nil {
			r.replaceTariffAltText(ctx, tariff.ID, dto)
		}

		if dto.Elements != nil {
			r.replaceElements(ctx, tariff.ID, dto)
		}

		return &tariff
	}

	return nil
}

func (r *TariffResolver) ReplaceTariffsByIdentifier(ctx context.Context, countryCode *string, partyID *string, cdrID *int64, dto []*TariffDto) {
	for _, tariffDto := range dto {
		r.ReplaceTariffByIdentifier(ctx, countryCode, partyID, *tariffDto.ID, cdrID, tariffDto)
	}
}

func (r *TariffResolver) replaceTariffAltText(ctx context.Context, tariffID int64, dto *TariffDto) {
	r.Repository.DeleteTariffAltTexts(ctx, tariffID)

	for _, displayTextDto := range dto.TariffAltText {
		displayTextParams := displaytext.NewCreateDisplayTextParams(displayTextDto)

		if displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams); err == nil {
			r.Repository.SetTariffAltText(ctx, db.SetTariffAltTextParams{
				TariffID:      tariffID,
				DisplayTextID: displayText.ID,
			})
		}
	}
}

func (r *TariffResolver) replaceElements(ctx context.Context, tariffID int64, dto *TariffDto) {
	r.ElementResolver.ReplaceElements(ctx, tariffID, dto.Elements)
}
