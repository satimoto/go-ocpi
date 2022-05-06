package sync

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func (r *SyncResolver) SynchronizeCredential(ctx context.Context, credential db.Credential) {
	if credential.VersionID.Valid {
		version, err := r.VersionResolver.Repository.GetVersion(ctx, credential.VersionID.Int64)

		if err != nil {
			util.LogOnError("OCPI006", "Error retrieving credential version", err)
			log.Printf("OCPI006: CredentialID=%v, VersionID=%v", credential.ID, credential.VersionID)
			return
		}

		if version.Version == "2.1.1" {
			r.Sync2_1_1Resolver.SynchronizeCredential(ctx, credential)
		}
	}
}
