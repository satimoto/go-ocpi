package sync

import (
	"context"
	"log"
	"sync"
	"time"
)

func (r *SyncService) StartService(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting Sync Service")
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup

	go r.startLoop()
}

func (r *SyncService) startLoop() {
	lastUpdated := time.Now().UTC().Add(time.Hour * -24)

	for {
		ctx := context.Background()
		
		if credentials, err := r.CredentialRepository.ListCredentials(ctx); err == nil {
			updatedDate := time.Now().UTC()

			for _, credential := range credentials {
				if credential.ClientToken.Valid {
					r.SynchronizeCredential(credential, &lastUpdated, nil, nil)
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
