package tariff

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/element"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TariffPullPayload struct {
	ID            *string                           `json:"id"`
	CountryCode   *string                           `json:"country_code"`
	PartyID       *string                           `json:"party_id"`
	Currency      *string                           `json:"currency"`
	TariffAltText []*displaytext.DisplayTextPayload `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string                           `json:"tariff_alt_url,omitempty"`
	Elements      []*element.ElementPayload         `json:"elements"`
	EnergyMix     *energymix.EnergyMixPayload       `json:"energy_mix,omitempty"`
	LastUpdated   *time.Time                        `json:"last_updated"`
}

type TariffPushPayload struct {
	ID            *string                           `json:"id"`
	Currency      *string                           `json:"currency"`
	TariffAltText []*displaytext.DisplayTextPayload `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string                           `json:"tariff_alt_url,omitempty"`
	Elements      []*element.ElementPayload         `json:"elements"`
	EnergyMix     *energymix.EnergyMixPayload       `json:"energy_mix,omitempty"`
	LastUpdated   *time.Time                        `json:"last_updated"`
}

func (r *TariffPushPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewTariffPushPayload(tariff db.Tariff) *TariffPushPayload {
	return &TariffPushPayload{
		ID:           &tariff.Uid,
		Currency:     &tariff.Currency,
		TariffAltUrl: util.NilString(tariff.TariffAltUrl.String),
		LastUpdated:  &tariff.LastUpdated,
	}
}

func NewCreateTariffParams(payload *TariffPushPayload) db.CreateTariffParams {
	return db.CreateTariffParams{
		Uid:          *payload.ID,
		Currency:     *payload.Currency,
		TariffAltUrl: util.SqlNullString(payload.TariffAltUrl),
		LastUpdated:  *payload.LastUpdated,
	}
}

func NewUpdateTariffByUidParams(tariff db.Tariff) db.UpdateTariffByUidParams {
	return db.UpdateTariffByUidParams{
		Uid:          tariff.Uid,
		Currency:     tariff.Currency,
		TariffAltUrl: tariff.TariffAltUrl,
		EnergyMixID:  tariff.EnergyMixID,
		LastUpdated:  tariff.LastUpdated,
	}
}

func (r *TariffResolver) CreateTariffPushPayload(ctx context.Context, tariff db.Tariff) *TariffPushPayload {
	response := NewTariffPushPayload(tariff)

	if tariffAltTexts, err := r.Repository.ListTariffAltTexts(ctx, tariff.ID); err == nil {
		response.TariffAltText = r.DisplayTextResolver.CreateDisplayTextListPayload(ctx, tariffAltTexts)
	}

	if elements, err := r.ElementResolver.Repository.ListElements(ctx, tariff.ID); err == nil {
		response.Elements = r.ElementResolver.CreateElementListPayload(ctx, elements)
	}

	if tariff.EnergyMixID.Valid {
		if energyMix, err := r.EnergyMixResolver.Repository.GetEnergyMix(ctx, tariff.EnergyMixID.Int64); err == nil {
			response.EnergyMix = r.EnergyMixResolver.CreateEnergyMixPayload(ctx, energyMix)
		}
	}

	return response
}

func (r *TariffResolver) CreateTariffListPayload(ctx context.Context, tariffs []db.Tariff) []*TariffPushPayload {
	list := []*TariffPushPayload{}
	for _, tariff := range tariffs {
		list = append(list, r.CreateTariffPushPayload(ctx, tariff))
	}
	return list
}
