package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/ocpirpc"
)

func NewCreateTokenResponse(token db.Token) *ocpirpc.CreateTokenResponse {
	return &ocpirpc.CreateTokenResponse{
		Id:        token.ID,
		Type:      string(token.Type),
		AuthId:    token.AuthID,
		Allowed:   string(token.Allowed),
		Whitelist: string(token.Whitelist),
	}
}

func NewUpdateTokensResponse(input ocpirpc.UpdateTokensRequest) *ocpirpc.UpdateTokensResponse {
	return &ocpirpc.UpdateTokensResponse{
		UserId:    input.UserId,
		Uid:       input.Uid,
		Allowed:   input.Allowed,
		Whitelist: input.Whitelist,
	}
}
