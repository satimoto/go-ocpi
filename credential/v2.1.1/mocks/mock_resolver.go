package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	businessdetail "github.com/satimoto/go-ocpi-api/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi-api/credential/v2.1.1"
	"github.com/satimoto/go-ocpi-api/util"
	version "github.com/satimoto/go-ocpi-api/version/mocks"
	versiondetail "github.com/satimoto/go-ocpi-api/version/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *credential.CredentialResolver {
	repo := credential.CredentialRepository(repositoryService)

	return &credential.CredentialResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		VersionResolver:        version.NewResolver(repositoryService, requester),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, requester),
	}
}
