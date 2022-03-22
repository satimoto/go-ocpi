package credential

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/satimoto/go-ocpi-api/internal/version"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

type CredentialRepository interface {
	GetCredentialByPartyAndCountryCode(ctx context.Context, arg db.GetCredentialByPartyAndCountryCodeParams) (db.Credential, error)
	GetCredentialByServerToken(ctx context.Context, serverToken sql.NullString) (db.Credential, error)
	UpdateCredential(ctx context.Context, arg db.UpdateCredentialParams) (db.Credential, error)
}

type CredentialResolver struct {
	Repository CredentialRepository
	*businessdetail.BusinessDetailResolver
	*version.VersionResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	repo := CredentialRepository(repositoryService)
	return &CredentialResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		VersionResolver:        version.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}

func (r *CredentialResolver) ReplaceCredential(ctx context.Context, credential db.Credential, payload *CredentialPayload) (*db.Credential, error) {
	if payload != nil {
		token := credential.ClientToken.String
		url := credential.Url
		partyID := credential.PartyID
		countryCode := credential.CountryCode

		if payload.Token != nil {
			token = *payload.Token
		}

		if payload.Url != nil {
			url = *payload.Url
		}

		if payload.PartyID != nil {
			partyID = *payload.PartyID
		}

		if payload.CountryCode != nil {
			countryCode = *payload.CountryCode
		}

		header := util.OCPIRequestHeader{
			Authentication: &token,
			ToCountryCode:  &countryCode,
			ToPartyId:      &partyID,
		}

		r.VersionResolver.UpdateVersions(ctx, url, header, credential.ID)
		version := r.VersionResolver.PickBestVersion(ctx, credential.ID)

		if version != nil {
			r.VersionDetailResolver.UpdateVersionDetail(ctx, version.Url, header, version.ID)

			params := db.UpdateCredentialParams{
				ID:          credential.ID,
				ClientToken: util.SqlNullString(token),
				ServerToken: util.SqlNullString(uuid.NewString()),
				Url:         url,
				PartyID:     partyID,
				CountryCode: countryCode,
				LastUpdated: time.Now(),
			}

			if payload.BusinessDetail != nil {
				r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, &credential.BusinessDetailID, payload.BusinessDetail)
			}

			if cred, err := r.Repository.UpdateCredential(ctx, params); err == nil {
				return &cred, nil
			}
		} else {
			return nil, ocpi.OCPIUnsupportedVersion(nil)
		}
	}

	return nil, ocpi.OCPIRegistrationError(nil)
}

func (r *CredentialResolver) updateCredential(ctx context.Context, cred db.Credential, rw http.ResponseWriter, request *http.Request) {
	payload := CredentialPayload{}

	if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
		render.Render(rw, request, ocpi.OCPIServerError(nil, err.Error()))
	}

	c, err := r.ReplaceCredential(ctx, cred, &payload)

	if err != nil {
		errResponse := err.(*ocpi.OCPIResponse)
		render.Render(rw, request, errResponse)
	}

	credentialPayload := r.CreateCredentialPayload(ctx, *c)
	render.Render(rw, request, ocpi.OCPISuccess(credentialPayload))
}
