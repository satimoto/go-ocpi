package credential

import "github.com/satimoto/go-datastore/pkg/db"

func NewUpdateCredentialParams(credential db.Credential) db.UpdateCredentialParams {
	return db.UpdateCredentialParams{
		ID:          credential.ID,
		ClientToken: credential.ClientToken,
		ServerToken: credential.ServerToken,
		Url:         credential.Url,
		CountryCode: credential.CountryCode,
		PartyID:     credential.PartyID,
		IsHub:       credential.IsHub,
		LastUpdated: credential.LastUpdated,
	}
}
