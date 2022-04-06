package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	"github.com/satimoto/go-ocpi-api/internal/version"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *ocpi.OCPIRequester) *version.VersionResolver {
	repo := version.VersionRepository(repositoryService)

	return &version.VersionResolver{
		Repository:    repo,
		OCPIRequester: requester,
	}
}
