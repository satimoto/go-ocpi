package tariff

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/element"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	"github.com/satimoto/go-ocpi-api/internal/tariffrestriction"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

type TariffRepository interface {
	CreateTariff(ctx context.Context, arg db.CreateTariffParams) (db.Tariff, error)
	DeleteTariffAltTexts(ctx context.Context, tariffID int64) error
	DeleteTariffByUid(ctx context.Context, uid string) error
	GetTariffByLastUpdated(ctx context.Context, arg db.GetTariffByLastUpdatedParams) (db.Tariff, error)
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
	*tariffrestriction.TariffRestrictionResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TariffResolver {
	repo := TariffRepository(repositoryService)
	return &TariffResolver{
		Repository:                repo,
		DisplayTextResolver:       displaytext.NewResolver(repositoryService),
		ElementResolver:           element.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		TariffRestrictionResolver: tariffrestriction.NewResolver(repositoryService),
		VersionDetailResolver:     versiondetail.NewResolver(repositoryService),
	}
}
