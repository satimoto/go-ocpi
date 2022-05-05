package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OCPIRequester) *versiondetail.VersionDetailResolver {
	repo := versiondetail.VersionDetailRepository(repositoryService)

	return &versiondetail.VersionDetailResolver{
		Repository:    repo,
		OCPIRequester: ocpiRequester,
	}
}
