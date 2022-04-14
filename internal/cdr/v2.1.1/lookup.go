package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func (r *CdrResolver) GetLastCdrByIdentity(ctx context.Context, countryCode string, partyID string) (db.Cdr, error) {
	params := db.GetCdrByIdentityOrderByLastUpdatedParams{
		CountryCode: util.SqlNullString(countryCode),
		PartyID:     util.SqlNullString(partyID),
	}

	return r.Repository.GetCdrByIdentityOrderByLastUpdated(ctx, params)
}
