package token

import (
	"context"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *TokenResolver) CreateTokenDto(ctx context.Context, token db.Token) *dto.TokenDto {
	return dto.NewTokenDto(token)
}

func (r *TokenResolver) CreateTokenListDto(ctx context.Context, tokens []db.Token) []render.Renderer {
	list := []render.Renderer{}
	
	for _, token := range tokens {
		list = append(list, r.CreateTokenDto(ctx, token))
	}

	return list
}
