package tariff

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/displaytext"
)

func (r *TariffResolver) ReplaceTariffByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, cdrID *int64, dto *TariffDto) *db.Tariff {
	if dto != nil {
		tariff, err := r.Repository.GetTariffByUid(ctx, uid)
		energyMixID := util.SqlNullInt64(tariff.EnergyMixID)
		tariffRestrictionID := util.SqlNullInt64(tariff.TariffRestrictionID)

		if dto.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, &energyMixID, dto.EnergyMix)
		}

		if dto.Restriction != nil {
			r.TariffRestrictionResolver.ReplaceTariffByIdentifierRestriction(ctx, &tariffRestrictionID, dto.Restriction)
		}

		if err == nil {
			tariffParams := param.NewUpdateTariffByUidParams(tariff)
			tariffParams.CountryCode = util.SqlNullString(countryCode)
			tariffParams.PartyID = util.SqlNullString(partyID)
			tariffParams.EnergyMixID = energyMixID
			tariffParams.TariffRestrictionID = tariffRestrictionID

			if dto.Currency != nil {
				tariffParams.Currency = *dto.Currency
			}

			if dto.LastUpdated != nil {
				tariffParams.LastUpdated = *dto.LastUpdated
			}

			if dto.TariffAltUrl != nil {
				tariffParams.TariffAltUrl = util.SqlNullString(dto.TariffAltUrl)
			}

			updatedTariff, err := r.Repository.UpdateTariffByUid(ctx, tariffParams)

			if err != nil {
				util.LogOnError("OCPI178", "Error updating tariff", err)
				log.Printf("OCPI178: Params=%#v", tariffParams)
				return nil
			}

			tariff = updatedTariff
		} else {
			tariffParams := NewCreateTariffParams(dto)
			tariffParams.CredentialID = credential.ID
			tariffParams.CountryCode = util.SqlNullString(countryCode)
			tariffParams.PartyID = util.SqlNullString(partyID)
			tariffParams.CdrID = util.SqlNullInt64(cdrID)
			tariffParams.EnergyMixID = energyMixID
			tariffParams.TariffRestrictionID = tariffRestrictionID

			tariff, err = r.Repository.CreateTariff(ctx, tariffParams)

			if err != nil {
				util.LogOnError("OCPI179", "Error creating tariff", err)
				log.Printf("OCPI179: Params=%#v", tariffParams)
				return nil
			}
		}

		if dto.TariffAltText != nil {
			r.replaceTariffAltText(ctx, tariff.ID, dto)
		}

		if dto.Elements != nil {
			r.replaceElements(ctx, tariff, dto)
		}

		return &tariff
	}

	return nil
}

func (r *TariffResolver) ReplaceTariffsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, cdrID *int64, dto []*TariffDto) {
	for _, tariffDto := range dto {
		r.ReplaceTariffByIdentifier(ctx, credential, countryCode, partyID, *tariffDto.ID, cdrID, tariffDto)
	}
}

func (r *TariffResolver) replaceTariffAltText(ctx context.Context, tariffID int64, dto *TariffDto) {
	r.Repository.DeleteTariffAltTexts(ctx, tariffID)

	for _, displayTextDto := range dto.TariffAltText {
		displayTextParams := displaytext.NewCreateDisplayTextParams(displayTextDto)
		displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams)

		if err != nil {
			util.LogOnError("OCPI180", "Error creating display text", err)
			log.Printf("OCPI180: Params=%#v", displayTextParams)
		}

		setTariffAltTextParams := db.SetTariffAltTextParams{
			TariffID:      tariffID,
			DisplayTextID: displayText.ID,
		}
		err = r.Repository.SetTariffAltText(ctx, setTariffAltTextParams)

		if err != nil {
			util.LogOnError("OCPI181", "Error setting tariff alt text", err)
			log.Printf("OCPI181: Params=%#v", setTariffAltTextParams)
		}
	}
}

func (r *TariffResolver) replaceElements(ctx context.Context, tariff db.Tariff, dto *TariffDto) {
	r.ElementResolver.ReplaceElements(ctx, tariff, dto.Elements)
}
