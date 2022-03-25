package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/satimoto/go-ocpi-api/internal/util"
)

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if dateFromParam := request.URL.Query().Get("date_from"); dateFromParam != "" {
			dateFrom := util.ParseTime(dateFromParam)

			if dateFrom != nil {
				ctx = context.WithValue(ctx, "date_from", dateFrom)
			}
		}

		if dateToParam := request.URL.Query().Get("date_to"); dateToParam != "" {
			dateTo := util.ParseTime(dateToParam)

			if dateTo != nil {
				ctx = context.WithValue(ctx, "date_to", dateTo)
			}
		}

		if offsetParam := request.URL.Query().Get("offset"); offsetParam != "" {
			offset, err := strconv.Atoi(offsetParam)

			if err == nil {
				ctx = context.WithValue(ctx, "offset", offset)
			}
		}

		if limitParam := request.URL.Query().Get("limit"); limitParam != "" {
			limit, err := strconv.Atoi(limitParam)

			if err == nil {
				ctx = context.WithValue(ctx, "limit", limit)
			}
		}

		next.ServeHTTP(rw, request.WithContext(ctx))
	})
}
