package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	businessdetail "github.com/satimoto/go-ocpi-api/internal/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi-api/internal/credential"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *credential.CredentialResolver {
	repo := credential.CredentialRepository(repositoryService)

	return &credential.CredentialResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
	}
}
