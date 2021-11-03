package rest

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type Repository interface {
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
}

type BaseRepository interface {}

type BeseResolver interface {}

type Resolver struct {
	Repository db.Repository
}
