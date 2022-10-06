package versiondetail

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/versiondetail"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type VersionDetailResolver struct {
	Repository    versiondetail.VersionDetailRepository
	OcpiRequester *transportation.OcpiRequester
}

func NewResolver(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *VersionDetailResolver {
	return &VersionDetailResolver{
		Repository:    versiondetail.NewRepository(repositoryService),
		OcpiRequester: ocpiRequester,
	}
}
