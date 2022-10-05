package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func NewCreateTokenParams(tokenDto *dto.TokenDto) db.CreateTokenParams {
	return db.CreateTokenParams{
		Uid:          *tokenDto.Uid,
		Type:         *tokenDto.Type,
		AuthID:       *tokenDto.AuthID,
		VisualNumber: util.SqlNullString(tokenDto.VisualNumber),
		Issuer:       *tokenDto.Issuer,
		Valid:        *tokenDto.Valid,
		Whitelist:    *tokenDto.Whitelist,
		Language:     util.SqlNullString(tokenDto.Language),
		LastUpdated:  *tokenDto.LastUpdated,
	}
}
