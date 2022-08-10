package credential

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewCredentialResponse(credential db.Credential) *ocpirpc.CreateCredentialResponse {
	return &ocpirpc.CreateCredentialResponse{
		Id:          credential.ID,
		ClientToken: credential.ClientToken.String,
		Url:         credential.Url,
		PartyId:     credential.PartyID,
		CountryCode: credential.CountryCode,
		IsHub:       credential.IsHub,
	}
}
