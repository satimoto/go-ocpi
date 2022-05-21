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
