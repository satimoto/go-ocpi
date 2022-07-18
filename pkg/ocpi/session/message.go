package session

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewSessionCreatedRequest(session db.Session) *ocpirpc.SessionCreatedRequest {
	return &ocpirpc.SessionCreatedRequest{
		SessionUid: session.Uid,
	}
}

func NewSessionUpdatedRequest(session db.Session) *ocpirpc.SessionUpdatedRequest {
	return &ocpirpc.SessionUpdatedRequest{
		SessionUid: session.Uid,
		Status: string(session.Status),
	}
}
