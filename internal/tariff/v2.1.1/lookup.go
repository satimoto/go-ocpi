package tariff

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func (r *TariffResolver) GetLastTariffByIdentity(ctx context.Context, countryCode string, partyID string) (db.Tariff, error) {
	params := db.GetTariffByIdentityOrderByLastUpdatedParams{
		CountryCode: util.SqlNullString(countryCode),
		PartyID:     util.SqlNullString(partyID),
	}

	return r.Repository.GetTariffByIdentityOrderByLastUpdated(ctx, params)
}
