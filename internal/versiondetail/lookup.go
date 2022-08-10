package versiondetail

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *VersionDetailResolver) GetVersionEndpointByIdentity(ctx context.Context, identifier string, countryCode string, partyID string) (db.VersionEndpoint, error) {
	return r.Repository.GetVersionEndpointByIdentity(ctx, db.GetVersionEndpointByIdentityParams{
		Identifier:  identifier,
		CountryCode: countryCode,
		PartyID:     partyID,
	})
}
