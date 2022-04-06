package version

import (
	"context"

	"github.com/satimoto/go-datastore/db"
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
