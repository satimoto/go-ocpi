package credential

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func GetCredentialByToken(r CredentialRepository, ctx context.Context, request *http.Request) (db.Credential, error) {
	if token := util.GetAuthenticationToken(request); token != "" {
		return r.GetCredentialByServerToken(ctx, util.SqlNullString(token))
	}

	return db.Credential{}, sql.ErrNoRows
}

func CredentialContextByToken(r CredentialRepository, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if credential, err := GetCredentialByToken(r, ctx, request); err == nil {
			ctx = context.WithValue(ctx, "credential", credential)
			next.ServeHTTP(rw, request.WithContext(ctx))
			return
		}

		render.Render(rw, request, ocpi.OCPIErrorUnknownResource(nil))
	})
}

func CredentialContextByPartyAndCountry(r CredentialRepository, next http.Handler) http.Handler {
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
				return
			}
		}

		render.Render(rw, request, ocpi.OCPIErrorUnknownResource(nil))
	})
}
