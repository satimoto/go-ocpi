package credentialrpc

import "github.com/satimoto/go-datastore/db"

func NewCredentialResponse(credential db.Credential) *CredentialResponse {
	return &CredentialResponse{
		Id:          credential.ID,
		ClientToken: credential.ClientToken.String,
		Url:         credential.Url,
		PartyId:     credential.PartyID,
		CountryCode: credential.CountryCode,
		IsHub:       credential.IsHub,
	}
}
