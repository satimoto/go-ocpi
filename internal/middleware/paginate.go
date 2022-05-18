package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/satimoto/go-datastore/pkg/util"
)

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if dateFromParam := request.URL.Query().Get("date_from"); dateFromParam != "" {
			dateFrom := util.ParseTime(dateFromParam, nil)

			if dateFrom != nil {
				ctx = context.WithValue(ctx, "date_from", dateFrom)
			}
		}

		if dateToParam := request.URL.Query().Get("date_to"); dateToParam != "" {
			dateTo := util.ParseTime(dateToParam, nil)

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

func GetDateFrom(ctx context.Context, defaultValue string) string {
	dateFrom := ctx.Value("date_from")

	if dateFrom != nil {
		return dateFrom.(string)
	}

	return defaultValue
}

func GetDateTo(ctx context.Context, defaultValue string) string {
	dateTo := ctx.Value("date_to")

	if dateTo != nil {
		return dateTo.(string)
	}

	return defaultValue
}

func GetLimit(ctx context.Context, defaultValue int64) int64 {
	limit := ctx.Value("limit")

	if limit != nil {
		return limit.(int64)
	}

	return defaultValue
}

func GetOffset(ctx context.Context, defaultValue int64) int64 {
	limit := ctx.Value("offset")

	if limit != nil {
		return limit.(int64)
	}

	return defaultValue
}
