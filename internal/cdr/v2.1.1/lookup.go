package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *CdrResolver) GetLastCdrByIdentity(ctx context.Context, credentialID *int64, countryCode *string, partyID *string) (db.Cdr, error) {
	params := db.GetCdrByLastUpdatedParams{
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

	return r.Repository.GetCdrByLastUpdated(ctx, params)
}
