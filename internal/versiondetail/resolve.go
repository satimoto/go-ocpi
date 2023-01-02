package versiondetail

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/versiondetail"
	"github.com/satimoto/go-ocpi/internal/service"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type VersionDetailResolver struct {
	Repository  versiondetail.VersionDetailRepository
	OcpiService *transportation.OcpiService
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *VersionDetailResolver {
	return &VersionDetailResolver{
		Repository:  versiondetail.NewRepository(repositoryService),
		OcpiService: services.OcpiService,
	}
}
