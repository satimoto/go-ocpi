package credential

import (
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewCreateCredentialParams(input ocpirpc.CreateCredentialRequest) db.CreateCredentialParams {
	return db.CreateCredentialParams{
		ClientToken: util.SqlNullString(input.ClientToken),
		Url:         input.Url,
		CountryCode: input.CountryCode,
		PartyID:     input.PartyId,
		IsHub:       input.IsHub,
		LastUpdated: time.Now(),
	}
}
