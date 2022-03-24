package tariff

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
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
	SetTariffAltText(ctx context.Context, arg db.SetTariffAltTextParams) error
	UpdateTariffByUid(ctx context.Context, arg db.UpdateTariffByUidParams) (db.Tariff, error)
}

type TariffResolver struct {
	Repository TariffRepository
	*credential.CredentialResolver
	*displaytext.DisplayTextResolver
	*element.ElementResolver
	*energymix.EnergyMixResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TariffResolver {
	repo := TariffRepository(repositoryService)
	return &TariffResolver{
		Repository:          repo,
		CredentialResolver:  credential.NewResolver(repositoryService),
		DisplayTextResolver: displaytext.NewResolver(repositoryService),
		ElementResolver:     element.NewResolver(repositoryService),
		EnergyMixResolver:   energymix.NewResolver(repositoryService),
	}
}

func (r *TariffResolver) ReplaceTariff(ctx context.Context, uid string, payload *TariffPayload) *db.Tariff {
	if payload != nil {
		tariff, err := r.Repository.GetTariffByUid(ctx, uid)
		energyMixID := util.NilInt64(tariff.EnergyMixID)

		if payload.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, energyMixID, payload.EnergyMix)
		}

		if err == nil {
			tariffParams := NewUpdateTariffByUidParams(tariff)
			tariffParams.EnergyMixID = util.SqlNullInt64(energyMixID)

			if payload.Currency != nil {
				tariffParams.Currency = *payload.Currency
			}

			if payload.LastUpdated != nil {
				tariffParams.LastUpdated = *payload.LastUpdated
			}

			if payload.TariffAltUrl != nil {
				tariffParams.TariffAltUrl = util.SqlNullString(payload.TariffAltUrl)
			}

			tariff, err = r.Repository.UpdateTariffByUid(ctx, tariffParams)
		} else {
			tariffParams := NewCreateTariffParams(payload)
			tariffParams.EnergyMixID = util.SqlNullInt64(energyMixID)

			tariff, err = r.Repository.CreateTariff(ctx, tariffParams)
		}

		if payload.TariffAltText != nil {
			r.replaceTariffAltText(ctx, tariff.ID, payload)
		}

		if payload.Elements != nil {
			r.replaceElements(ctx, tariff.ID, payload)
		}

		return &tariff
	}

	return nil
}

func (r *TariffResolver) replaceTariffAltText(ctx context.Context, tariffID int64, payload *TariffPayload) {
	r.Repository.DeleteTariffAltTexts(ctx, tariffID)

	for _, displayTextPayload := range payload.TariffAltText {
		displayTextParams := displaytext.NewCreateDisplayTextParams(displayTextPayload)

		if displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams); err == nil {
			r.Repository.SetTariffAltText(ctx, db.SetTariffAltTextParams{
				TariffID:      tariffID,
				DisplayTextID: displayText.ID,
			})
		}
	}
}

func (r *TariffResolver) replaceElements(ctx context.Context, tariffID int64, payload *TariffPayload) {
	r.ElementResolver.ReplaceElements(ctx, tariffID, payload.Elements)
}
