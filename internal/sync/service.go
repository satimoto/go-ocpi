package sync

import (
	"context"
	"log"
	"sync"
	"time"
)

func (r *SyncService) StartService(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting Sync Service")
	r.HtbService.StartService(shutdownCtx, waitGroup)
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup


	go r.startLoop()
}

func (r *SyncService) startLoop() {
	lastUpdated := time.Now().UTC().Add(time.Hour * -1)

	for {
		ctx := context.Background()
		
		if credentials, err := r.CredentialRepository.ListCredentials(ctx); err == nil {
			updatedDate := time.Now().UTC()

			for _, credential := range credentials {
				if credential.ClientToken.Valid {
					if credential.CountryCode == "FR" && credential.PartyID == "007" {
						r.HtbService.Run(ctx, credential, lastUpdated)
					}

					r.SynchronizeCredential(credential, false, &lastUpdated, nil, nil)
				}
			}

			lastUpdated = updatedDate
		}

		select {
		case <-r.shutdownCtx.Done():
			log.Printf("Shutting down Sync Service")
			return
		case <-time.After(time.Hour * 24):
			continue
		}
	}
}
