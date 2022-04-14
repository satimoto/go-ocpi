package location

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func (r *LocationResolver) GetLastLocationByIdentity(ctx context.Context, countryCode string, partyID string) (db.Location, error) {
	params := db.GetLocationByIdentityOrderByLastUpdatedParams{
		CountryCode: util.SqlNullString(countryCode),
		PartyID:     util.SqlNullString(partyID),
	}

	return r.Repository.GetLocationByIdentityOrderByLastUpdated(ctx, params)
}
