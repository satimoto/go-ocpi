package sync

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *SyncResolver) SynchronizeCredential(ctx context.Context, credential db.Credential) {
	if credential.ClientToken.Valid {
		r.LocationResolver.PullLocationsByIdentifier(ctx, credential.CountryCode, credential.PartyID, credential.ClientToken.String)
		r.TariffResolver.PullTariffsByIdentifier(ctx, credential.CountryCode, credential.PartyID, credential.ClientToken.String)
		r.SessionResolver.PullSessionsByIdentifier(ctx, credential.CountryCode, credential.PartyID, credential.ClientToken.String)
		r.CdrResolver.PullCdrsByIdentifier(ctx, credential.CountryCode, credential.PartyID, credential.ClientToken.String)
	}
}
