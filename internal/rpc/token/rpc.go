package token

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi-api/ocpirpc/tokenrpc"
)

func (r *RpcTokenResolver) CreateToken(ctx context.Context, request *tokenrpc.CreateTokenRequest) (*tokenrpc.CreateTokenResponse, error) {
	if request != nil {
		dto := NewCreateTokenDto(request)
		tokenAllowed := db.TokenAllowedTypeALLOWED
		authID, err := r.TokenResolver.GenerateAuthID(ctx)

		if err != nil {
			log.Printf("Error CreateToken GenerateAuthID: %v", err)
			log.Printf("Request=%#v", request)
			return nil, errors.New("Error creating token")
		}

		dto.AuthID = &authID
		dto.VisualNumber = &authID
		dto.Issuer = util.NilString(os.Getenv("ISSUER"))

		if len(request.Allowed) > 0 {
			tokenAllowed = db.TokenAllowedType(request.Allowed)
		}

		if len(request.Whitelist) == 0 {
			dto.Whitelist = NilTokenWhitelistType(db.TokenWhitelistTypeALLOWED)
		}

		token := r.TokenResolver.ReplaceToken(ctx, request.UserId, tokenAllowed, *dto.Uid, dto)

		if token == nil {
			log.Printf("Error CreateToken ReplaceToken: %v", err)
			log.Printf("Dto=%#v", dto)
			return nil, errors.New("Error creating token")
		}

		return tokenrpc.NewCreateTokenResponse(*token), nil
	}

	return nil, errors.New("Error creating token")
}

func (r *RpcTokenResolver) UpdateTokens(ctx context.Context, request *tokenrpc.UpdateTokensRequest) (*tokenrpc.UpdateTokensResponse, error) {
	if request != nil {
		tokens, err := r.TokenResolver.Repository.ListTokensByUserID(ctx, request.UserId)

		if err != nil {
			log.Printf("Error UpdateTokens ListTokensByUserID: %v", err)
			log.Printf("Request=%#v", request)
			return nil, errors.New("Error updating tokens")
		}

		for _, t := range tokens {
			if len(request.Uid) == 0 || request.Uid == t.Uid {
				dto := &token.TokenDto{
					Uid: &request.Uid, 
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

		return tokenrpc.NewUpdateTokensResponse(*request), nil
	}

	return nil, errors.New("Error updating tokens")
}
