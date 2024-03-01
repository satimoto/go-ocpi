package sync

import (
	"context"
	"log"
	"sync"
	"time"
)

func (r *SyncService) StartService(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting Sync Service")
	r.AtCpiService.StartService(shutdownCtx, waitGroup)
	r.DeIseService.StartService(shutdownCtx, waitGroup)
	r.ItBecService.StartService(shutdownCtx, waitGroup)
	r.NlConService.StartService(shutdownCtx, waitGroup)
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup

	go r.startLoop()
}

func (r *SyncService) startLoop() {
	timeNow := time.Now().UTC()
	lastUpdated := time.Date(
		timeNow.Year(),
		timeNow.Month(),
		timeNow.Day(),
		0,
		10,
		0,
		0,
		timeNow.Location())
	startTime := lastUpdated.Add(time.Hour * 24)

	if lastUpdated.After(timeNow) {
		startTime = lastUpdated
		lastUpdated = startTime.Add(time.Hour * -24)
	}

	waitDuration := startTime.Sub(timeNow)

	for {
		select {
		case <-r.shutdownCtx.Done():
			log.Printf("Shutting down Sync Service")
			return
		case <-time.After(waitDuration):
		}

		waitDuration = time.Hour * 24
		ctx := context.Background()

		if credentials, err := r.CredentialRepository.ListCredentials(ctx); err == nil {
			updatedDate := time.Now().UTC()

			for _, credential := range credentials {
				if credential.ClientToken.Valid {
					if credential.CountryCode == "FR" && credential.PartyID == "007" {
						r.AtCpiService.Run(ctx, credential, lastUpdated)
						r.DeIseService.Run(ctx, credential, lastUpdated)
						r.ItBecService.Run(ctx, credential, lastUpdated)
						r.NlConService.Run(ctx, credential, lastUpdated)
					}

					r.SynchronizeCredential(credential, false, true, &lastUpdated, nil, nil)
				}
			}

			lastUpdated = updatedDate
		}
	}
}
