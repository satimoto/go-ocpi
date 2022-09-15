package credential

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewCreateCredentialParams(input ocpirpc.CreateCredentialRequest) db.CreateCredentialParams {
	return db.CreateCredentialParams{
		ClientToken: dbUtil.SqlNullString(input.ClientToken),
		Url:         input.Url,
		CountryCode: input.CountryCode,
		PartyID:     input.PartyId,
		IsHub:       input.IsHub,
		LastUpdated: util.NewTimeUTC(),
	}
}
