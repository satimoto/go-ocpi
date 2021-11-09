package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	version "github.com/satimoto/go-ocpi-api/version/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *version.VersionDetailResolver {
	repo := version.VersionDetailRepository(repositoryService)

	return &version.VersionDetailResolver{
		Repository:             repo,
	}
}
