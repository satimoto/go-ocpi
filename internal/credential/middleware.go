package credential

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-datastore/util"
)

func GetCredentialByToken(r CredentialRepository, ctx context.Context, request *http.Request) (db.Credential, error) {
	if token := GetAuthenticationToken(request); token != "" {
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

		render.Render(rw, request, transportation.OCPIErrorUnknownResource(nil))
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

		render.Render(rw, request, transportation.OCPIErrorUnknownResource(nil))
	})
}

func GetCredential(ctx context.Context) *db.Credential {
	credential := ctx.Value("credential")

	if credential != nil {
		return credential.(*db.Credential)
	}

	return nil
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