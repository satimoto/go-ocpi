package credential

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *CredentialResolver) ReplaceCredential(ctx context.Context, credential db.Credential, dto *CredentialDto) (*db.Credential, error) {
	if dto != nil {
		token := credential.ClientToken.String
		url := credential.Url
		partyID := credential.PartyID
		countryCode := credential.CountryCode

		if dto.Token != nil {
			token = *dto.Token
		}

		if dto.Url != nil {
			url = *dto.Url
		}

		if dto.PartyID != nil {
			partyID = *dto.PartyID
		}

		if dto.CountryCode != nil {
			countryCode = *dto.CountryCode
		}

		header := transportation.NewOCPIRequestHeader(&token, nil, nil)

		r.VersionResolver.PullVersions(ctx, url, header, credential.ID)
		version := r.VersionResolver.GetBestVersion(ctx, credential.ID)

		if version != nil {
			r.VersionDetailResolver.PullVersionEndpoints(ctx, version.Url, header, version.ID)

			params := db.UpdateCredentialParams{
				ID:          credential.ID,
				ClientToken: util.SqlNullString(token),
				ServerToken: util.SqlNullString(uuid.NewString()),
				Url:         url,
				PartyID:     partyID,
				CountryCode: countryCode,
				LastUpdated: time.Now(),
			}

			if dto.BusinessDetail != nil {
				r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, &credential.BusinessDetailID, dto.BusinessDetail)
			}

			if cred, err := r.Repository.UpdateCredential(ctx, params); err == nil {
				go r.SyncResolver.SynchronizeCredential(ctx, cred)

				return &cred, nil
			}
		} else {
			return nil, transportation.OCPIUnsupportedVersion(nil)
		}
	}

	return nil, transportation.OCPIRegistrationError(nil)
}
