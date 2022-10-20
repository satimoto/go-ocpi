package token

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/middleware"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

const (
	X_LIMIT = 1000
)

func (r *TokenResolver) ListTokens(rw http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	filterDateFrom := middleware.GetDateFrom(ctx, "")
	filterDateTo := middleware.GetDateTo(ctx, "")
	filterLimit := middleware.GetLimit(ctx, X_LIMIT)
	filterOffset := middleware.GetOffset(ctx, 0)

	if filterLimit > X_LIMIT {
		filterLimit = X_LIMIT
	}

	listTokensParams := db.ListTokensParams{
		FilterDateFrom: filterDateFrom,
		FilterDateTo:   filterDateTo,
		FilterLimit:    filterLimit,
		FilterOffset:   filterOffset,
	}

	tokens, err := r.Repository.ListTokens(ctx, listTokensParams)

	if err != nil {
		util.LogOnError("OCPI204", "Error listing tokens", err)
		log.Printf("OCPI204: Params=%#v", listTokensParams)

		render.Render(rw, request, transportation.OcpiServerError(nil, err.Error()))
		return
	}

	dto := r.CreateTokenListDto(ctx, tokens)

	countTokensParams := db.CountTokensParams{
		FilterDateFrom: middleware.GetDateFrom(ctx, ""),
		FilterDateTo:   middleware.GetDateTo(ctx, ""),
	}

	if count, err := r.Repository.CountTokens(ctx, countTokensParams); err == nil {
		nextOffset := filterOffset + filterLimit

		if len(tokens) == X_LIMIT || nextOffset < count {
			requestUrl := request.URL
			query := requestUrl.Query()
			query.Set("limit", strconv.FormatInt(filterLimit, 10))
			query.Set("offset", strconv.FormatInt(nextOffset, 10))

			unescapedQuery, _ := url.QueryUnescape(query.Encode())
			requestUrl.RawQuery = unescapedQuery

			rw.Header().Add("Link", fmt.Sprintf("<%s%s>; rel=\"next\"", os.Getenv("API_DOMAIN"), requestUrl.String()))
		}

		rw.Header().Add("X-Limit", strconv.FormatInt(filterLimit, 10))
		rw.Header().Add("X-Total-Count", strconv.FormatInt(count, 10))
	}

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
