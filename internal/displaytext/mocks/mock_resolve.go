package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	displaytextMocks "github.com/satimoto/go-datastore/pkg/displaytext/mocks"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *displaytext.DisplayTextResolver {
	return &displaytext.DisplayTextResolver{
		Repository: displaytextMocks.NewRepository(repositoryService),
	}
}
