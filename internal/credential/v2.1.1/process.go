package credential

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
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

		if dto.CountryCode != nil {
			countryCode = *dto.CountryCode
		}

		if dto.PartyID != nil {
			partyID = *dto.PartyID
		}

		return r.CredentialResolver.RegisterCredential(ctx, credential, token, url, countryCode, partyID)
	}

	return nil, transportation.OcpiRegistrationError(nil)
}
