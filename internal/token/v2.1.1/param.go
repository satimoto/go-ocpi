package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateTokenParams(dto *TokenDto) db.CreateTokenParams {
	return db.CreateTokenParams{
		Uid:          *dto.Uid,
		Type:         *dto.Type,
		AuthID:       *dto.AuthID,
		VisualNumber: util.SqlNullString(dto.VisualNumber),
		Issuer:       *dto.Issuer,
		Valid:        *dto.Valid,
		Whitelist:    *dto.Whitelist,
		Language:     util.SqlNullString(dto.Language),
		LastUpdated:  *dto.LastUpdated,
	}
}

func NewUpdateTokenByUidParams(token db.Token) db.UpdateTokenByUidParams {
	return db.UpdateTokenByUidParams{
		Uid:          token.Uid,
		Type:         token.Type,
		AuthID:       token.AuthID,
		VisualNumber: token.VisualNumber,
		Issuer:       token.Issuer,
		Allowed:      token.Allowed,
		Valid:        token.Valid,
		Whitelist:    token.Whitelist,
		Language:     token.Language,
		LastUpdated:  token.LastUpdated,
	}
}
