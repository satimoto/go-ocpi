package token

import (
	"math"
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/middleware"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *TokenResolver) ListTokens(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	tokens, err := r.Repository.ListTokens(ctx, db.ListTokensParams{
		FilterDateFrom: middleware.GetDateFrom(ctx, ""),
		FilterDateTo:   middleware.GetDateTo(ctx, ""),
		FilterLimit:    middleware.GetLimit(ctx, math.MaxInt64),
		FilterOffset:   middleware.GetOffset(ctx, 0),
	})

	if err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	dto := r.CreateTokenListDto(ctx, tokens)

	if err := render.Render(rw, request, ocpi.OCPISuccess(dto)); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}
}

func (r *TokenResolver) AuthorizeToken(rw http.ResponseWriter, request *http.Request) {
	r.TokenAuthorizationResolver.AuthorizeToken(rw, request)
}
