package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	partyMocks "github.com/satimoto/go-datastore/pkg/party/mocks"
	"github.com/satimoto/go-ocpi/internal/party"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *party.PartyResolver {
	return &party.PartyResolver{
		Repository: partyMocks.NewRepository(repositoryService),
	}
}
