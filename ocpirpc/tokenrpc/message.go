package tokenrpc

import "github.com/satimoto/go-datastore/db"

func NewTokenResponse(token db.Token) *TokenResponse {
	return &TokenResponse{
		Id:        token.ID,
		Type:      string(token.Type),
		AuthId:    token.AuthID,
		Allowed:   string(token.Allowed),
		Whitelist: string(token.Whitelist),
	}
}
