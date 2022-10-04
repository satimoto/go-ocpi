package sync

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *SyncService) SynchronizeCredential(credential db.Credential, lastUpdated *time.Time, countryCode *string, partyID *string) {
	if credential.VersionID.Valid {
		ctx := context.Background()

		log.Printf("Sync credential Url=%v LastUpdated=%v CountryCode=%v PartyID=%v", credential.Url, lastUpdated, countryCode, partyID)
		version, err := r.VersionResolver.Repository.GetVersion(ctx, credential.VersionID.Int64)

		if err != nil {
			util.LogOnError("OCPI270", "Error retrieving credential version", err)
			log.Printf("OCPI270: CredentialID=%v, VersionID=%v", credential.ID, credential.VersionID)
			return
		}

		for _, syncerHandler := range r.syncerHandlers {
			if syncerHandler.Version == version.Version {
				syncerHandler.Syncer.SyncByIdentifier(ctx, credential, lastUpdated, countryCode, partyID)
			}
		}
	}
}
