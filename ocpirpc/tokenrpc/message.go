package tokenrpc

import "github.com/satimoto/go-datastore/db"

func NewCreateTokenResponse(token db.Token) *CreateTokenResponse {
	return &CreateTokenResponse{
		Id:        token.ID,
		Type:      string(token.Type),
		AuthId:    token.AuthID,
		Allowed:   string(token.Allowed),
		Whitelist: string(token.Whitelist),
	}
}

func NewUpdateTokensResponse(input UpdateTokensRequest) *UpdateTokensResponse {
	return &UpdateTokensResponse{
		UserId: input.UserId,
		Uid: input.Uid,
		Allowed: input.Allowed,
		Whitelist: input.Whitelist,
	}
}
