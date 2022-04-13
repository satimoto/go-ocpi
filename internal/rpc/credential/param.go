package credential

import (
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
	"github.com/satimoto/go-ocpi-api/ocpirpc/credentialrpc"
)

func NewCreateBusinessDetailParams(input ocpirpc.BusinessDetail) db.CreateBusinessDetailParams {
	return db.CreateBusinessDetailParams{
		Name:    input.Name,
		Website: util.SqlNullString(input.Website),
	}
}

func NewCreateCredentialParams(input credentialrpc.CreateCredentialRequest) db.CreateCredentialParams {
	return db.CreateCredentialParams{
		ClientToken: util.SqlNullString(input.ClientToken),
		Url:         input.Url,
		CountryCode: input.CountryCode,
		PartyID:     input.PartyId,
		IsHub:       input.IsHub,
		LastUpdated: time.Now(),
	}
}

func NewCreateImageParams(input ocpirpc.Image) db.CreateImageParams {
	return db.CreateImageParams{
		Url:       input.Url,
		Thumbnail: util.SqlNullString(input.Thumbnail),
		Category:  db.ImageCategory(input.Category),
		Width:     util.SqlNullInt32(input.Width),
		Height:    util.SqlNullInt32(input.Height),
	}
}
