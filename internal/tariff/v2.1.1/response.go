package tariff

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/element"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	"github.com/satimoto/go-ocpi-api/internal/tariffrestriction"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TariffPullDto struct {
	ID            *string                                 `json:"id"`
	CountryCode   *string                                 `json:"country_code"`
	PartyID       *string                                 `json:"party_id"`
	Currency      *string                                 `json:"currency"`
	TariffAltText []*displaytext.DisplayTextDto           `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string                                 `json:"tariff_alt_url,omitempty"`
	Elements      []*element.ElementDto                   `json:"elements"`
	EnergyMix     *energymix.EnergyMixDto                 `json:"energy_mix,omitempty"`
	Restriction   *tariffrestriction.TariffRestrictionDto `json:"restriction,omitempty"`
	LastUpdated   *time.Time                              `json:"last_updated"`
}

type TariffPushDto struct {
	ID            *string                                 `json:"id"`
	Currency      *string                                 `json:"currency"`
	TariffAltText []*displaytext.DisplayTextDto           `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string                                 `json:"tariff_alt_url,omitempty"`
	Elements      []*element.ElementDto                   `json:"elements"`
	EnergyMix     *energymix.EnergyMixDto                 `json:"energy_mix,omitempty"`
	Restriction   *tariffrestriction.TariffRestrictionDto `json:"restriction,omitempty"`
	LastUpdated   *time.Time                              `json:"last_updated"`
}

func (r *TariffPushDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewTariffPushDto(tariff db.Tariff) *TariffPushDto {
	return &TariffPushDto{
		ID:           &tariff.Uid,
		Currency:     &tariff.Currency,
		TariffAltUrl: util.NilString(tariff.TariffAltUrl.String),
		LastUpdated:  &tariff.LastUpdated,
	}
}

func NewCreateTariffParams(dto *TariffPushDto) db.CreateTariffParams {
	return db.CreateTariffParams{
		Uid:          *dto.ID,
		Currency:     *dto.Currency,
		TariffAltUrl: util.SqlNullString(dto.TariffAltUrl),
		LastUpdated:  *dto.LastUpdated,
	}
}

func NewUpdateTariffByUidParams(tariff db.Tariff) db.UpdateTariffByUidParams {
	return db.UpdateTariffByUidParams{
		Uid:                 tariff.Uid,
		CountryCode:         tariff.CountryCode,
		PartyID:             tariff.PartyID,
		Currency:            tariff.Currency,
		TariffAltUrl:        tariff.TariffAltUrl,
		EnergyMixID:         tariff.EnergyMixID,
		TariffRestrictionID: tariff.TariffRestrictionID,
		LastUpdated:         tariff.LastUpdated,
	}
}

func (r *TariffResolver) CreateTariffPushDto(ctx context.Context, tariff db.Tariff) *TariffPushDto {
	response := NewTariffPushDto(tariff)

	if tariffAltTexts, err := r.Repository.ListTariffAltTexts(ctx, tariff.ID); err == nil {
		response.TariffAltText = r.DisplayTextResolver.CreateDisplayTextListDto(ctx, tariffAltTexts)
	}

	if elements, err := r.ElementResolver.Repository.ListElements(ctx, tariff.ID); err == nil {
		response.Elements = r.ElementResolver.CreateElementListDto(ctx, elements)
	}

	if tariff.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixResolver.Repository.GetEnergyMix(ctx, tariff.EnergyMixID.Int64); err == nil {
			response.EnergyMix = r.EnergyMixResolver.CreateEnergyMixDto(ctx, energyMix)
		}
	}

	if tariff.TariffRestrictionID.Valid {
		if tariffRestriction, err := r.TariffRestrictionResolver.Repository.GetTariffRestriction(ctx, tariff.TariffRestrictionID.Int64); err == nil {
			response.Restriction = r.TariffRestrictionResolver.CreateTariffRestrictionDto(ctx, tariffRestriction)
		}
	}

	return response
}

func (r *TariffResolver) CreateTariffPushListDto(ctx context.Context, tariffs []db.Tariff) []*TariffPushDto {
	list := []*TariffPushDto{}
	for _, tariff := range tariffs {
		list = append(list, r.CreateTariffPushDto(ctx, tariff))
	}
	return list
}
