package versiondetail

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/versiondetail"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

type VersionDetailResolver struct {
	Repository    versiondetail.VersionDetailRepository
	OcpiRequester *transportation.OcpiRequester
}

func NewResolver(repositoryService *db.RepositoryService) *VersionDetailResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *VersionDetailResolver {
	return &VersionDetailResolver{
		Repository:    versiondetail.NewRepository(repositoryService),
		OcpiRequester: ocpiRequester,
	}
}
