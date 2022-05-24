package cdr

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewCdrCreatedRequest(cdr db.Cdr) *ocpirpc.CdrCreatedRequest {
	return &ocpirpc.CdrCreatedRequest{
		CdrUid: cdr.Uid,
	}
}
