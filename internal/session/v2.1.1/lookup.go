package session

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *SessionResolver) GetLastSessionByIdentity(ctx context.Context, credentialID *int64, countryCode *string, partyID *string) (db.Session, error) {
	params := db.GetSessionByLastUpdatedParams{
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

	return r.Repository.GetSessionByLastUpdated(ctx, params)
}
