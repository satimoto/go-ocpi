package sync

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	coreTariff "github.com/satimoto/go-ocpi/internal/tariff"
)

func (r *SyncService) SynchronizeCredential(credential db.Credential, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	if credential.VersionID.Valid {
		activeSyncsKey := fmt.Sprintf("%s*%s", credential.CountryCode, credential.PartyID)

		if countryCode != nil && partyID != nil {
			activeSyncsKey = fmt.Sprintf("%s*%s", *countryCode, *partyID)
		}

		if _, ok := r.activeSyncs[activeSyncsKey]; !ok {
			r.activeSyncs[activeSyncsKey] = true
			ctx := context.Background()

			log.Printf("Start credential sync %v Url=%v LastUpdated=%v CountryCode=%v PartyID=%v",
				activeSyncsKey, credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))
			version, err := r.VersionResolver.Repository.GetVersion(ctx, credential.VersionID.Int64)

			if err != nil {
				metrics.RecordError("OCPI270", "Error retrieving credential version", err)
				log.Printf("OCPI270: CredentialID=%v, VersionID=%v", credential.ID, credential.VersionID)
				return
			}

			for _, syncerHandler := range r.syncerHandlers {
				if syncerHandler.Version == version.Version {
					if countryCode == nil || syncerHandler.Identifier == coreLocation.IDENTIFIER || syncerHandler.Identifier == coreTariff.IDENTIFIER {
						if syncerHandler.Identifier == coreTariff.IDENTIFIER && !r.tariffsSyncing {
							// Only sync tariffs one at a time
							r.tariffsSyncing = true
							syncerHandler.Syncer.SyncByIdentifier(ctx, credential, fullSync, lastUpdated, countryCode, partyID)
							r.tariffsSyncing = false
						} else if syncerHandler.Identifier != coreTariff.IDENTIFIER {
							syncerHandler.Syncer.SyncByIdentifier(ctx, credential, fullSync, lastUpdated, countryCode, partyID)
						}
					}
				}
			}

			delete(r.activeSyncs, activeSyncsKey)
			log.Printf("End credential sync %v", activeSyncsKey)
			return
		}

		log.Printf("Ignore credential sync %v Url=%v LastUpdated=%v CountryCode=%v PartyID=%v",
			activeSyncsKey, credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))
	}
}
