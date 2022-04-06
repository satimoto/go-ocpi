package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *displaytext.DisplayTextResolver {
	repo := displaytext.DisplayTextRepository(repositoryService)

	return &displaytext.DisplayTextResolver{
		Repository: repo,
	}
}
