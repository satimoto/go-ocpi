package cdr

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/async"
	coreCommand "github.com/satimoto/go-ocpi/internal/command"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *CdrResolver) updateCommand(cdr db.Cdr, session db.Session) {
	ctx := context.Background()

	if cdr.AuthorizationID.Valid {
		updateCommandStopBySessionIDParams := db.UpdateCommandStopBySessionIDParams{
			SessionID:   session.Uid,
			Status:      db.CommandResponseTypeACCEPTED,
			LastUpdated: time.Now().UTC(),
		}

		command, err := r.CommandRepository.UpdateCommandStopBySessionID(ctx, updateCommandStopBySessionIDParams)

		if err != nil {
			metrics.RecordError("OCPI329", "Error updating command stop", err)
			log.Printf("OCPI329: SessionUid=%#v", session.Uid)
			return
		}

		asyncKey := fmt.Sprintf(coreCommand.STOP_COMMAND_ASYNC_KEY, command.ID)
		asyncResult := async.AsyncResult{
			String: string(command.Status),
			Bool:   command.Status == db.CommandResponseTypeACCEPTED,
		}

		r.AsyncService.Set(asyncKey, asyncResult)
	}
}
