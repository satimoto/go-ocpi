package tariff

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *TariffResolver) ReplaceTariffByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, cdrID *int64, tariffDto *dto.TariffDto) *db.Tariff {
	if tariffDto != nil {
		tariff, err := r.Repository.GetTariffByUid(ctx, uid)
		energyMixID := util.SqlNullZeroInt64(tariff.EnergyMixID)
		tariffRestrictionID := util.SqlNullZeroInt64(tariff.TariffRestrictionID)

		if tariffDto.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, &energyMixID, tariffDto.EnergyMix)
		}

		if tariffDto.Restriction != nil {
			r.TariffRestrictionResolver.ReplaceTariffByIdentifierRestriction(ctx, &tariffRestrictionID, tariffDto.Restriction)
		}

		if err == nil && cdrID == nil {
			tariffParams := param.NewUpdateTariffByUidParams(tariff)
			tariffParams.CountryCode = util.SqlNullString(countryCode)
			tariffParams.PartyID = util.SqlNullString(partyID)
			tariffParams.EnergyMixID = energyMixID
			tariffParams.TariffRestrictionID = tariffRestrictionID

			if tariffDto.CountryCode != nil {
				tariffParams.CountryCode = util.SqlNullString(tariffDto.CountryCode)
			}

			if tariffDto.PartyID != nil {
				tariffParams.PartyID = util.SqlNullString(tariffDto.PartyID)
			}

			if tariffDto.Currency != nil {
				tariffParams.Currency = *tariffDto.Currency
			}

			if tariffDto.LastUpdated != nil {
				tariffParams.LastUpdated = tariffDto.LastUpdated.Time()
			}

			if tariffDto.TariffAltUrl != nil {
				tariffParams.TariffAltUrl = util.SqlNullString(tariffDto.TariffAltUrl)
			}

			updatedTariff, err := r.Repository.UpdateTariffByUid(ctx, tariffParams)

			if err != nil {
				metrics.RecordError("OCPI178", "Error updating tariff", err)
				log.Printf("OCPI178: Params=%#v", tariffParams)
				return nil
			}

			tariff = updatedTariff
		} else {
			tariffParams := NewCreateTariffParams(tariffDto)
			tariffParams.CredentialID = credential.ID
			tariffParams.CdrID = util.SqlNullInt64(cdrID)
			tariffParams.EnergyMixID = energyMixID
			tariffParams.TariffRestrictionID = tariffRestrictionID

			if !tariffParams.CountryCode.Valid {
				tariffParams.CountryCode = util.SqlNullString(countryCode)
			}

			if !tariffParams.PartyID.Valid {
				tariffParams.PartyID = util.SqlNullString(partyID)
			}

			tariff, err = r.Repository.CreateTariff(ctx, tariffParams)

			if err != nil {
				metrics.RecordError("OCPI179", "Error creating tariff", err)
				log.Printf("OCPI179: Params=%#v", tariffParams)
				return nil
			}
		}

		if tariffDto.TariffAltText != nil {
			r.replaceTariffAltText(ctx, tariff.ID, tariffDto)
		}

		if tariffDto.Elements != nil {
			r.replaceElements(ctx, tariff, tariffDto)
		}

		return &tariff
	}

	return nil
}

func (r *TariffResolver) ReplaceTariffsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, cdrID *int64, tariffsDto []*dto.TariffDto) {
	for _, tariffDto := range tariffsDto {
		r.ReplaceTariffByIdentifier(ctx, credential, countryCode, partyID, *tariffDto.ID, cdrID, tariffDto)
	}
}

func (r *TariffResolver) replaceTariffAltText(ctx context.Context, tariffID int64, tariffDto *dto.TariffDto) {
	r.Repository.DeleteTariffAltTexts(ctx, tariffID)

	for _, displayTextDto := range tariffDto.TariffAltText {
		displayTextParams := displaytext.NewCreateDisplayTextParams(displayTextDto)
		displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams)

		if err != nil {
			metrics.RecordError("OCPI180", "Error creating display text", err)
			log.Printf("OCPI180: Params=%#v", displayTextParams)
		}

		setTariffAltTextParams := db.SetTariffAltTextParams{
			TariffID:      tariffID,
			DisplayTextID: displayText.ID,
		}
		err = r.Repository.SetTariffAltText(ctx, setTariffAltTextParams)

		if err != nil {
			metrics.RecordError("OCPI181", "Error setting tariff alt text", err)
			log.Printf("OCPI181: Params=%#v", setTariffAltTextParams)
		}
	}
}

func (r *TariffResolver) replaceElements(ctx context.Context, tariff db.Tariff, tariffDto *dto.TariffDto) {
	r.ElementResolver.ReplaceElements(ctx, tariff, tariffDto.Elements)
}
