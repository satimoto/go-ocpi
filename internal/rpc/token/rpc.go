package token

import (
	"context"
	"errors"
	"os"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/ocpirpc/tokenrpc"
)

func (r *RpcCredentialResolver) CreateToken(ctx context.Context, in *tokenrpc.CreateTokenRequest) (*tokenrpc.TokenResponse, error) {
	if in != nil {
		params := tokenrpc.NewCreateTokenParams(*in)
		authID, err := r.TokenResolver.GenerateAuthID(ctx)

		if err != nil {
			return nil, err
		}

		params.AuthID = authID
		params.VisualNumber = util.SqlNullString(authID)
		params.Issuer = os.Getenv("ISSUER")

		if len(in.Allowed) == 0 {
			params.Allowed = db.TokenAllowedTypeALLOWED
		}

		if len(in.Whitelist) == 0 {
			params.Whitelist = db.TokenWhitelistTypeALLOWED
		}

		t, err := r.TokenResolver.Repository.CreateToken(ctx, params)

		if err != nil {
			return nil, err
		}

		return tokenrpc.NewTokenResponse(t), nil
	}

	return nil, errors.New("Missing CreateTokenRequest")
}
