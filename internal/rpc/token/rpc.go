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
			log.Printf("Error CreateToken GenerateAuthID: %v", err)
			log.Printf("Request=%#v", request)
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
			log.Printf("Error CreateToken ReplaceToken: %v", err)
			log.Printf("Dto=%#v", dto)
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
			log.Printf("Error UpdateTokens ListTokensByUserID: %v", err)
			log.Printf("Request=%#v", request)
			return nil, errors.New("error updating tokens")
		}

		for _, t := range tokens {
			if len(request.Uid) == 0 || request.Uid == t.Uid {
				dto := &token.TokenDto{
					Uid:         &request.Uid,
					LastUpdated: util.NilTime(time.Now()),
				}

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
