package credentialrpc

import (
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreateCredentialParams(input CreateCredentialRequest) db.CreateCredentialParams {
	return db.CreateCredentialParams{
		ClientToken: util.SqlNullString(input.ClientToken),
		Url:         input.Url,
		CountryCode: input.CountryCode,
		PartyID:     input.PartyId,
		IsHub:       input.IsHub,
		LastUpdated: time.Now(),
	}
}
