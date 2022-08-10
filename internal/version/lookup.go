package version

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *VersionResolver) GetBestVersion(ctx context.Context, credentialID int64) *db.Version {
	if versions, err := r.Repository.ListVersions(ctx, credentialID); err == nil {
		for _, version := range versions {
			if version.Version == "2.1.1" {
				return &version
			}
		}
	}

	return nil
}

func (r *VersionResolver) GetPreferredVersion(ctx context.Context, credentialID int64, version string) *db.Version {
	if versions, err := r.Repository.ListVersions(ctx, credentialID); err == nil {
		for _, v := range versions {
			if v.Version == version {
				return &v
			}
		}
	}

	return nil
}