package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1/mocks"
	displaytext "github.com/satimoto/go-ocpi-api/internal/displaytext/mocks"
	element "github.com/satimoto/go-ocpi-api/internal/element/mocks"
	energymix "github.com/satimoto/go-ocpi-api/internal/energymix/mocks"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *tariff.TariffResolver {
	repo := tariff.TariffRepository(repositoryService)

	return &tariff.TariffResolver{
		Repository:             repo,
		CredentialResolver:  credential.NewResolver(repositoryService, requester),
		DisplayTextResolver: displaytext.NewResolver(repositoryService),
		ElementResolver:     element.NewResolver(repositoryService),
		EnergyMixResolver:   energymix.NewResolver(repositoryService),
	}
}
