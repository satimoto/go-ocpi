package sync

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *SyncResolver) SynchronizeCredential(ctx context.Context, credential db.Credential) {
	if credential.VersionID.Valid {
		version, err := r.VersionResolver.Repository.GetVersion(ctx, credential.VersionID.Int64)

		if err != nil {
			util.LogOnError("OCPI270", "Error retrieving credential version", err)
			log.Printf("OCPI270: CredentialID=%v, VersionID=%v", credential.ID, credential.VersionID)
			return
		}

		if version.Version == "2.1.1" {
			r.SyncResolver_2_1_1.SynchronizeCredential(ctx, credential)
		}
	}
}
