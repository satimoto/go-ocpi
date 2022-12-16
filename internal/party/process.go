package party

import (
	"context"
	"errors"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *PartyResolver) GetParty(ctx context.Context, credential db.Credential, countryCode, partyID *string) (*db.Party, error) {
	if countryCode != nil && partyID != nil {
		getPartyByCredentialParams := db.GetPartyByCredentialParams{
			CredentialID: credential.ID,
			CountryCode:  *countryCode,
			PartyID:      *partyID,
		}

		if party, err := r.Repository.GetPartyByCredential(ctx, getPartyByCredentialParams); err == nil {
			return &party, nil
		}
	}

	getPartyByCredentialParams := db.GetPartyByCredentialParams{
		CredentialID: credential.ID,
		CountryCode:  credential.CountryCode,
		PartyID:      credential.PartyID,
	}

	if credentialParty, err := r.Repository.GetPartyByCredential(ctx, getPartyByCredentialParams); err == nil {
		if countryCode != nil && partyID != nil {
			createPartyParams := db.CreatePartyParams{
				CredentialID:             credential.ID,
				CountryCode:              *countryCode,
				PartyID:                  *partyID,
				IsIntermediateCdrCapable: credentialParty.IsIntermediateCdrCapable,
				PublishLocation:          credentialParty.PublishLocation,
				PublishNullTariff:        credentialParty.PublishNullTariff,
			}
	
			if party, err := r.Repository.CreateParty(ctx, createPartyParams); err == nil {
				return &party, nil
			}
		}
	
		return &credentialParty, err
	}

	return nil, errors.New("Error getting party")
}
