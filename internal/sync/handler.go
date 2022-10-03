package sync

import (
	"context"
	"log"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
)

type Syncer interface {
	SyncByIdentifier(ctx context.Context, credential db.Credential, lastUpdated *time.Time, countryCode *string, partyID *string)
}

type SyncerHandler struct {
	Version string
	Syncer  Syncer
}

func (r *SyncService) AddHandler(version string, syncer Syncer) {
	log.Printf("Added sync handler for %v", version)

	r.syncerHandlers = append(r.syncerHandlers, &SyncerHandler{
		Version: version,
		Syncer:  syncer,
	})
}
