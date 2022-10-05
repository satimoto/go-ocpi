package dto

import (
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type OcpiTariffsDto struct {
	Data          []*TariffDto  `json:"data,omitempty"`
	StatusCode    int16         `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	Timestamp     ocpitype.Time `json:"timestamp"`
}

type TariffDto struct {
	ID            *string                       `json:"id"`
	CountryCode   *string                       `json:"country_code,omitempty"`
	PartyID       *string                       `json:"party_id,omitempty"`
	Currency      *string                       `json:"currency"`
	TariffAltText []*coreDto.DisplayTextDto     `json:"tariff_alt_text,omitempty"`
	TariffAltUrl  *string                       `json:"tariff_alt_url,omitempty"`
	Elements      []*coreDto.ElementDto         `json:"elements"`
	EnergyMix     *coreDto.EnergyMixDto         `json:"energy_mix,omitempty"`
	Restriction   *coreDto.TariffRestrictionDto `json:"restriction,omitempty"`
	LastUpdated   *time.Time                    `json:"last_updated"`
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
