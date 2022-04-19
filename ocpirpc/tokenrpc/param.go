package tokenrpc

import (
	"time"

	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/db"
)

func NewCreateTokenParams(input CreateTokenRequest) db.CreateTokenParams {
	return db.CreateTokenParams{
		Uid:         uuid.NewString(),
		UserID:      input.UserId,
		Type:        db.TokenType(input.Type),
		Allowed:     db.TokenAllowedType(input.Allowed),
		Valid:       true,
		Whitelist:   db.TokenWhitelistType(input.Whitelist),
		LastUpdated: time.Now(),
	}
}
