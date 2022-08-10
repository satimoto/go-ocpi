package token

import (
	"log"
	"math"
	"net/http"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TokenResolver) ListTokens(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	listTokensParams := db.ListTokensParams{
		FilterDateFrom: middleware.GetDateFrom(ctx, ""),
		FilterDateTo:   middleware.GetDateTo(ctx, ""),
		FilterLimit:    middleware.GetLimit(ctx, math.MaxInt64),
		FilterOffset:   middleware.GetOffset(ctx, 0),
	}
	tokens, err := r.Repository.ListTokens(ctx, listTokensParams)

	if err != nil {
		util.LogOnError("OCPI204", "Error listing tokens", err)
		log.Printf("OCPI204: Params=%#v", listTokensParams)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	dto := r.CreateTokenListDto(ctx, tokens)

	if err := render.Render(rw, request, transportation.OcpiSuccess(dto)); err != nil {
		util.LogOnError("OCPI205", "Error rendering response", err)
		util.LogHttpRequest("OCPI205", request.URL.String(), request, true)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}
}

func (r *TokenResolver) AuthorizeToken(rw http.ResponseWriter, request *http.Request) {
	r.TokenAuthorizationResolver.AuthorizeToken(rw, request)
}
