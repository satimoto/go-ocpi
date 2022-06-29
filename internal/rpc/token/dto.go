package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewCreateTokenDto(input *ocpirpc.CreateTokenRequest) *token.TokenDto {
	return &token.TokenDto{
		Uid:         util.NilString(uuid.NewString()),
		Type:        NilTokenType(input.Type),
		Valid:       util.NilBool(true),
		Whitelist:   NilTokenWhitelistType(input.Whitelist),
		LastUpdated: util.NilTime(time.Now()),
	}
}

func NilTokenType(i interface{}) *db.TokenType {
	switch t := i.(type) {
	case db.TokenType:
		return &t
	case string:
		if len(t) > 0 {
			typed := db.TokenType(t)
			return &typed
		}
	}

	return nil
}

func NilTokenAllowedType(i interface{}) *db.TokenAllowedType {
	switch t := i.(type) {
	case db.TokenAllowedType:
		return &t
	case string:
		if len(t) > 0 {
			typed := db.TokenAllowedType(t)
			return &typed
		}
	}

	return nil
}

func NilTokenWhitelistType(i interface{}) *db.TokenWhitelistType {
	switch t := i.(type) {
	case db.TokenWhitelistType:
		return &t
	case string:
		if len(t) > 0 {
			typed := db.TokenWhitelistType(t)
			return &typed
		}
	}

	return nil
}
