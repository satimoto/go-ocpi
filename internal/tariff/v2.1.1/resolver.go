package tariff

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/element"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TariffRepository interface {
	CreateTariff(ctx context.Context, arg db.CreateTariffParams) (db.Tariff, error)
	DeleteTariffAltTexts(ctx context.Context, tariffID int64) error
	DeleteTariffByUid(ctx context.Context, uid string) error
	GetTariffByUid(ctx context.Context, uid string) (db.Tariff, error)
	ListTariffAltTexts(ctx context.Context, tariffID int64) ([]db.DisplayText, error)
	ListTariffsByCdr(ctx context.Context, cdrID sql.NullInt64) ([]db.Tariff, error)
	SetTariffAltText(ctx context.Context, arg db.SetTariffAltTextParams) error
	UpdateTariffByUid(ctx context.Context, arg db.UpdateTariffByUidParams) (db.Tariff, error)
}

type TariffResolver struct {
	Repository TariffRepository
	*displaytext.DisplayTextResolver
	*element.ElementResolver
	*energymix.EnergyMixResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TariffResolver {
	repo := TariffRepository(repositoryService)
	return &TariffResolver{
		Repository:          repo,
		DisplayTextResolver: displaytext.NewResolver(repositoryService),
		ElementResolver:     element.NewResolver(repositoryService),
		EnergyMixResolver:   energymix.NewResolver(repositoryService),
	}
}

func (r *TariffResolver) ReplaceTariff(ctx context.Context, countryCode *string, partyID *string, uid string, cdrID *int64, dto *TariffPushDto) *db.Tariff {
	if dto != nil {
		tariff, err := r.Repository.GetTariffByUid(ctx, uid)
		energyMixID := util.NilInt64(tariff.EnergyMixID)

		if dto.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, energyMixID, dto.EnergyMix)
		}

		if err == nil {
			tariffParams := NewUpdateTariffByUidParams(tariff)
			tariffParams.CountryCode = util.SqlNullString(countryCode)
			tariffParams.PartyID = util.SqlNullString(partyID)
			tariffParams.EnergyMixID = util.SqlNullInt64(energyMixID)

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

func (r *TariffResolver) replaceTariffAltText(ctx context.Context, tariffID int64, dto *TariffPushDto) {
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

func (r *TariffResolver) replaceElements(ctx context.Context, tariffID int64, dto *TariffPushDto) {
	r.ElementResolver.ReplaceElements(ctx, tariffID, dto.Elements)
}
