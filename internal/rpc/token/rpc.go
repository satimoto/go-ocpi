package token

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/ocpirpc"
	ocpiToken "github.com/satimoto/go-ocpi/pkg/ocpi/token"
)

func (r *RpcTokenResolver) CreateToken(ctx context.Context, request *ocpirpc.CreateTokenRequest) (*ocpirpc.CreateTokenResponse, error) {
	if request != nil {
		dto := NewCreateTokenDto(request)
		tokenAllowed := db.TokenAllowedTypeNOCREDIT
		authID, err := r.TokenResolver.GenerateAuthID(ctx)

		if err != nil {
			util.LogOnError("OCPI279", "Error generating AuthID", err)
			log.Printf("OCPI279: Request=%#v", request)
			return nil, errors.New("error creating token")
		}

		dto.AuthID = &authID
		dto.VisualNumber = &authID
		dto.Issuer = util.NilString(os.Getenv("ISSUER"))

		if len(request.Allowed) > 0 {
			tokenAllowed = db.TokenAllowedType(request.Allowed)
		}

		if len(request.Whitelist) == 0 {
			dto.Whitelist = NilTokenWhitelistType(db.TokenWhitelistTypeNEVER)
		}

		t := r.TokenResolver.ReplaceToken(ctx, request.UserId, tokenAllowed, *dto.Uid, dto)

		if t == nil {
			util.LogOnError("OCPI280", "Error replacing token", err)
			log.Printf("OCPI280: Dto=%#v", dto)
			return nil, errors.New("error creating token")
		}

		return ocpiToken.NewCreateTokenResponse(*t), nil
	}

	return nil, errors.New("error creating token")
}

func (r *RpcTokenResolver) UpdateTokens(ctx context.Context, request *ocpirpc.UpdateTokensRequest) (*ocpirpc.UpdateTokensResponse, error) {
	if request != nil {
		tokens, err := r.TokenResolver.Repository.ListTokensByUserID(ctx, request.UserId)

		if err != nil {
			util.LogOnError("OCPI281", "Error listing tokens", err)
			log.Printf("OCPI281: Request=%#v", request)
			return nil, errors.New("error updating tokens")
		}

		for _, t := range tokens {
			if len(request.Uid) == 0 || request.Uid == t.Uid {
				dto := token.NewTokenDto(t)
				dto.LastUpdated = util.NilTime(time.Now())

				tokenAllowed := t.Allowed

				if len(request.Allowed) > 0 {
					tokenAllowed = db.TokenAllowedType(request.Allowed)
				}

				if len(request.Whitelist) > 0 {
					dto.Whitelist = NilTokenWhitelistType(request.Whitelist)
				}

				r.TokenResolver.ReplaceToken(ctx, request.UserId, tokenAllowed, *dto.Uid, dto)
			}
		}

		return ocpiToken.NewUpdateTokensResponse(*request), nil
	}

	return nil, errors.New("error updating tokens")
}
