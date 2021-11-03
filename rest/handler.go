package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/util"
)

func Paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if dateFromParam := chi.URLParam(request, "date_from"); dateFromParam != "" {
			dateFrom := util.ParseTime(dateFromParam)

			if dateFrom != nil {
				ctx = context.WithValue(ctx, "dateFrom", dateFrom)
			}
		}

		if dateToParam := chi.URLParam(request, "date_to"); dateToParam != "" {
			dateTo := util.ParseTime(dateToParam)

			if dateTo != nil {
				ctx = context.WithValue(ctx, "dateTo", dateTo)
			}
		}

		if offsetParam := chi.URLParam(request, "offset"); offsetParam != "" {
			offset, err := strconv.Atoi(offsetParam)

			if err == nil {
				ctx = context.WithValue(ctx, "offset", offset)
			}
		}

		if limitParam := chi.URLParam(request, "limit"); limitParam != "" {
			limit, err := strconv.Atoi(limitParam)

			if err == nil {
				ctx = context.WithValue(ctx, "limit", limit)
			}
		}

		next.ServeHTTP(rw, request.WithContext(ctx))
	})
}

func Credential(r Repository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
			ctx := request.Context()

			partyId := chi.URLParam(request, "party_id")
			countryCode := chi.URLParam(request, "country_code")

			if partyId != "" && countryCode != "" {
				credential, err := r.GetCredentialByPartyAndCountryCode(ctx, db.GetCredentialByPartyAndCountryCodeParams{
					PartyID:     partyId,
					CountryCode: countryCode,
				})

				if err == nil {
					ctx = context.WithValue(ctx, "credential", credential)

					next.ServeHTTP(rw, request.WithContext(ctx))
				}
			}

			render.Render(rw, request, OCPIErrorNotEnoughInformation(nil))
		})
	}
}
