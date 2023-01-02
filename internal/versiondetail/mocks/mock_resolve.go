package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	versiondetailMocks "github.com/satimoto/go-datastore/pkg/versiondetail/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *versiondetail.VersionDetailResolver {
	return &versiondetail.VersionDetailResolver{
		Repository:  versiondetailMocks.NewRepository(repositoryService),
		OcpiService: services.OcpiService,
	}
}
