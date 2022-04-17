package tariff

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *TariffResolver) GetLastTariffByIdentity(ctx context.Context, credentialID *int64, countryCode *string, partyID *string) (db.Tariff, error) {
	params := db.GetTariffByLastUpdatedParams{
		CredentalID: -1,
	}

	if credentialID != nil {
		params.CredentalID = *credentialID
	}

	if countryCode != nil {
		params.CountryCode = *countryCode
	}
	if partyID != nil {
		params.PartyID = *partyID
	}

	return r.Repository.GetTariffByLastUpdated(ctx, params)
}
