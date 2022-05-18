package tariff

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/element"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	"github.com/satimoto/go-ocpi-api/internal/tariffrestriction"
)

type OcpiTariffsDto struct {
	Data          []*TariffDto `json:"data,omitempty"`
	StatusCode    int16        `json:"status_code"`
	StatusMessage string       `json:"status_message"`
	Timestamp     time.Time    `json:"timestamp"`
}

type TariffDto struct {
	ID            *string                                 `json:"id"`
	CountryCode   *string                                 `json:"country_code,omitempty"`
	PartyID       *string                                 `json:"party_id,omitempty"`
	Currency      *string                                 `json:"currency"`
	TariffAltText []*displaytext.DisplayTextDto           `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string                                 `json:"tariff_alt_url,omitempty"`
	Elements      []*element.ElementDto                   `json:"elements"`
	EnergyMix     *energymix.EnergyMixDto                 `json:"energy_mix,omitempty"`
	Restriction   *tariffrestriction.TariffRestrictionDto `json:"restriction,omitempty"`
	LastUpdated   *time.Time                              `json:"last_updated"`
}

func (r *TariffDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewTariffDto(tariff db.Tariff) *TariffDto {
	return &TariffDto{
		ID:           &tariff.Uid,
		Currency:     &tariff.Currency,
		TariffAltUrl: util.NilString(tariff.TariffAltUrl),
		LastUpdated:  &tariff.LastUpdated,
	}
}

func (r *TariffResolver) CreateTariffDto(ctx context.Context, tariff db.Tariff) *TariffDto {
	response := NewTariffDto(tariff)

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

func (r *TariffResolver) CreateTariffPushListDto(ctx context.Context, tariffs []db.Tariff) []*TariffDto {
	list := []*TariffDto{}
	for _, tariff := range tariffs {
		list = append(list, r.CreateTariffDto(ctx, tariff))
	}
	return list
}
