package token

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
	"github.com/satimoto/go-ocpi/internal/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
	ocpiToken "github.com/satimoto/go-ocpi/pkg/ocpi/token"
)

func (r *RpcTokenResolver) CreateToken(ctx context.Context, request *ocpirpc.CreateTokenRequest) (*ocpirpc.CreateTokenResponse, error) {
	if request != nil {
		tokenDto := NewCreateTokenDto(request)
		tokenAllowed := db.TokenAllowedTypeNOCREDIT
		authID, err := r.TokenResolver.GenerateAuthID(ctx)

		if err != nil {
			dbUtil.LogOnError("OCPI279", "Error generating AuthID", err)
			log.Printf("OCPI279: Request=%#v", request)
			return nil, errors.New("error creating token")
		}

		tokenDto.AuthID = &authID
		tokenDto.VisualNumber = &authID
		tokenDto.Issuer = dbUtil.NilString(os.Getenv("ISSUER"))

		if len(request.Allowed) > 0 {
			tokenAllowed = db.TokenAllowedType(request.Allowed)
		} else {
			getTokenByUserIDParams := db.GetTokenByUserIDParams{
				UserID: request.UserId,
				Type: db.TokenTypeOTHER,
			}
	
			if t, err := r.TokenResolver.Repository.GetTokenByUserID(ctx, getTokenByUserIDParams); err == nil {
				tokenAllowed = t.Allowed
			}	
		}

		if len(request.Whitelist) == 0 {
			tokenDto.Whitelist = NilTokenWhitelistType(db.TokenWhitelistTypeNEVER)
		}

		t := r.TokenResolver.ReplaceToken(ctx, request.UserId, tokenAllowed, *tokenDto.Uid, tokenDto)

		if t == nil {
			dbUtil.LogOnError("OCPI280", "Error replacing token", err)
			log.Printf("OCPI280: Dto=%#v", tokenDto)
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
			dbUtil.LogOnError("OCPI281", "Error listing tokens", err)
			log.Printf("OCPI281: Request=%#v", request)
			return nil, errors.New("error updating tokens")
		}

		for _, t := range tokens {
			if len(request.Uid) == 0 || request.Uid == t.Uid {
				lastUpdated := util.NewTimeUTC()
				tokenDto := dto.NewTokenDto(t)
				tokenDto.LastUpdated = ocpitype.NilOcpiTime(&lastUpdated)

				tokenAllowed := t.Allowed

				if len(request.Allowed) > 0 {
					tokenAllowed = db.TokenAllowedType(request.Allowed)
				}

				if len(request.Whitelist) > 0 {
					tokenDto.Whitelist = NilTokenWhitelistType(request.Whitelist)
				}

				r.TokenResolver.ReplaceToken(ctx, request.UserId, tokenAllowed, *tokenDto.Uid, tokenDto)
			}
		}

		return ocpiToken.NewUpdateTokensResponse(*request), nil
	}

	return nil, errors.New("error updating tokens")
}
