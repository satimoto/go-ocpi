package sync

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreCdr "github.com/satimoto/go-ocpi/internal/cdr"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	coreSession "github.com/satimoto/go-ocpi/internal/session"
	coreTariff "github.com/satimoto/go-ocpi/internal/tariff"
)

func (r *SyncService) SynchronizeCredential(credential db.Credential, fullSync, withTariffs bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	if credential.VersionID.Valid {
		log.Printf("Start credential sync Url=%v LastUpdated=%v CountryCode=%v PartyID=%v", credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))

		ctx := context.Background()
		version, err := r.VersionResolver.Repository.GetVersion(ctx, credential.VersionID.Int64)

		if err != nil {
			metrics.RecordError("OCPI270", "Error retrieving credential version", err)
			log.Printf("OCPI270: CredentialID=%v, VersionID=%v", credential.ID, credential.VersionID)
			return
		}

		if countryCode != nil && partyID != nil {
			// Single party sync
			r.synchronizeParty(ctx, credential, version, fullSync, lastUpdated, *countryCode, *partyID)
		} else {
			// Full sync all parties
			parties, _ := r.PartyRepository.ListPartiesByCredentialID(ctx, credential.ID)

			for _, party := range parties {
				r.synchronizeParty(ctx, credential, version, fullSync, lastUpdated, party.CountryCode, party.PartyID)
			}

			r.synchronizeSessionsAndCdrs(ctx, credential, version, fullSync, lastUpdated, countryCode, partyID)
		}

		if withTariffs {
			r.synchronizeTariffs(ctx, credential, version, fullSync, lastUpdated, countryCode, partyID)
		}
	}
}

func (r *SyncService) synchronizeParty(ctx context.Context, credential db.Credential, version db.Version, fullSync bool, lastUpdated *time.Time, countryCode, partyID string) {
	activeSyncsKey := fmt.Sprintf("%s*%s", countryCode, partyID)

	if _, ok := r.activeSyncs[activeSyncsKey]; !ok {
		r.activeSyncs[activeSyncsKey] = true

		log.Printf("Start party sync %v Url=%v LastUpdated=%v CountryCode=%v PartyID=%v", activeSyncsKey, credential.Url, lastUpdated, countryCode, partyID)

		for _, syncerHandler := range r.syncerHandlers {
			if syncerHandler.Version == version.Version {
				if syncerHandler.Identifier == coreLocation.IDENTIFIER {
					syncerHandler.Syncer.SyncByIdentifier(ctx, credential, fullSync, lastUpdated, &countryCode, &partyID)
				}
			}
		}

		delete(r.activeSyncs, activeSyncsKey)
		log.Printf("End party sync %v", activeSyncsKey)
	}
}

func (r *SyncService) synchronizeSessionsAndCdrs(ctx context.Context, credential db.Credential, version db.Version, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	for _, syncerHandler := range r.syncerHandlers {
		if syncerHandler.Version == version.Version && (syncerHandler.Identifier == coreSession.IDENTIFIER || syncerHandler.Identifier == coreCdr.IDENTIFIER) {
			syncerHandler.Syncer.SyncByIdentifier(ctx, credential, fullSync, lastUpdated, countryCode, partyID)
		}
	}
}

func (r *SyncService) synchronizeTariffs(ctx context.Context, credential db.Credential, version db.Version, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	if !r.tariffsSyncing {
		for _, syncerHandler := range r.syncerHandlers {
			if syncerHandler.Version == version.Version && syncerHandler.Identifier == coreTariff.IDENTIFIER {
				r.tariffsSyncing = true
				syncerHandler.Syncer.SyncByIdentifier(ctx, credential, fullSync, lastUpdated, countryCode, partyID)
				r.tariffsSyncing = false
			}
		}
	}
}
