package middleware

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type CredentialRepository interface {
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
}

func GetCredentialByToken(r CredentialRepository, ctx context.Context, request *http.Request) (db.Credential, error) {
	if token := GetAuthorizationToken(request); token != "" {
		return r.GetCredentialByServerToken(ctx, util.SqlNullString(token))
	}

	return db.Credential{}, sql.ErrNoRows
}

func CredentialContextByToken(r CredentialRepository, next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		if credential, err := GetCredentialByToken(r, ctx, request); err == nil {
			ctx = context.WithValue(ctx, "credential", &credential)
			next.ServeHTTP(rw, request.WithContext(ctx))
			return
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
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
				ctx = context.WithValue(ctx, "credential", &credential)
				next.ServeHTTP(rw, request.WithContext(ctx))
				return
			}
		}

		render.Render(rw, request, transportation.OcpiErrorUnknownResource(nil))
	})
}

func GetCredential(ctx context.Context) *db.Credential {
	credential := ctx.Value("credential")

	if credential != nil {
		return credential.(*db.Credential)
	}

	return nil
}

func GetAuthorizationToken(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	if len(authorization) > 6 && strings.ToUpper(authorization[0:5]) == "TOKEN" {
		return authorization[6:]
	}

	return ""
}

func GetCountryCode(request *http.Request) *string {
	ctx := request.Context()
	countryCode := chi.URLParam(request, "country_code")

	if countryCode == "" {
		if ctxCredential := ctx.Value("credential"); ctxCredential != nil {
			credential := ctxCredential.(*db.Credential)
			return &credential.CountryCode
		}
	}

	return &countryCode
}

func GetPartyID(request *http.Request) *string {
	ctx := request.Context()
	partyID := chi.URLParam(request, "party_id")

	if partyID == "" {
		if ctxCredential := ctx.Value("credential"); ctxCredential != nil {
			credential := ctxCredential.(*db.Credential)
			return &credential.PartyID
		}
	}

	return &partyID
}
