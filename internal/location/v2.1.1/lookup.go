package location

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *LocationResolver) GetLastLocationByIdentity(ctx context.Context, credentialID *int64, countryCode *string, partyID *string) (db.Location, error) {
	params := db.GetLocationByLastUpdatedParams{
		CredentialID: -1,
	}

	if credentialID != nil {
		params.CredentialID = *credentialID
	}

	if countryCode != nil {
		params.CountryCode = *countryCode
	}

	if partyID != nil {
		params.PartyID = *partyID
	}

	return r.Repository.GetLocationByLastUpdated(ctx, params)
}
