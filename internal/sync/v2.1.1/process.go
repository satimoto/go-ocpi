package sync

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *SyncResolver) SynchronizeCredential(ctx context.Context, credential db.Credential) {
	if credential.ClientToken.Valid {
		r.LocationResolver.PullLocationsByIdentifier(ctx, credential, nil, nil)
		r.TariffResolver.PullTariffsByIdentifier(ctx, credential, nil, nil)
		r.SessionResolver.PullSessionsByIdentifier(ctx, credential, nil, nil)
		r.CdrResolver.PullCdrsByIdentifier(ctx, credential, nil, nil)
	}
}
