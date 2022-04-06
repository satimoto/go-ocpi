package token

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TokenDto struct {
	Uid          *string                `json:"uid"`
	Type         *db.TokenType          `json:"type"`
	AuthID       *string                `json:"auth_id"`
	VisualNumber *string                `json:"visual_number,omitempty"`
	Issuer       *string                `json:"issuer"`
	Valid        *bool                  `json:"valid"`
	Whitelist    *db.TokenWhitelistType `json:"whitelist"`
	Language     *string                `json:"language,omitempty"`
	LastUpdated  *time.Time             `json:"last_updated"`
}

func (r *TokenDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewTokenDto(token db.Token) *TokenDto {
	return &TokenDto{
		Uid:          &token.Uid,
		Type:         &token.Type,
		AuthID:       &token.AuthID,
		VisualNumber: util.NilString(token.VisualNumber),
		Issuer:       &token.Issuer,
		Valid:        &token.Valid,
		Whitelist:    &token.Whitelist,
		Language:     util.NilString(token.Language),
		LastUpdated:  &token.LastUpdated,
	}
}

func (r *TokenResolver) CreateTokenDto(ctx context.Context, token db.Token) *TokenDto {
	return NewTokenDto(token)
}

func (r *TokenResolver) CreateTokenListDto(ctx context.Context, tokens []db.Token) []render.Renderer {
	list := []render.Renderer{}
	for _, token := range tokens {
		list = append(list, r.CreateTokenDto(ctx, token))
	}
	return list
}
