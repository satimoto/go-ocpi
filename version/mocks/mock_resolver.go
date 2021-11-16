package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/util"
	"github.com/satimoto/go-ocpi-api/version"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *version.VersionResolver {
	repo := version.VersionRepository(repositoryService)

	return &version.VersionResolver{
		Repository:         repo,
		OCPIRequester:      requester,
	}
}
