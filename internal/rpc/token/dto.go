package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
	"github.com/satimoto/go-ocpi/internal/util"
	"github.com/satimoto/go-ocpi/ocpirpc"
)

func NewCreateTokenDto(input *ocpirpc.CreateTokenRequest) *dto.TokenDto {
	lastUpdated := util.NewTimeUTC()

	return &dto.TokenDto{
		Uid:         NilUid(input.Uid),
		Type:        NilTokenType(input.Type),
		Valid:       dbUtil.NilBool(true),
		Whitelist:   NilTokenWhitelistType(input.Whitelist),
		LastUpdated: ocpitype.NilOcpiTime(&lastUpdated),
	}
}

func NilUid(uid string) *string {
	if len(uid) > 0 {
		return &uid
	}

	return nil
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
