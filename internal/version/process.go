package version

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *VersionResolver) ReplaceVersions(ctx context.Context, credentialID int64, dto []*VersionDto) []*db.Version {
	versions := []*db.Version{}

	if dto != nil {
		r.Repository.DeleteVersions(ctx, credentialID)

		for _, versionDto := range dto {
			versionParams := NewCreateVersionParams(credentialID, versionDto)

			if version, err := r.Repository.CreateVersion(ctx, versionParams); err == nil {
				versions = append(versions, &version)
			}
		}
	}

	return versions
}

func (r *VersionResolver) PullVersions(ctx context.Context, url string, header transportation.OcpiRequestHeader, credentialID int64) []*db.Version {
	if response, err := r.OcpiRequester.Do(http.MethodGet, url, header, nil); err == nil {
		ocpiResponse, err := r.UnmarshalPullDto(response.Body)
		response.Body.Close()

		if err == nil && ocpiResponse.StatusCode == transportation.STATUS_CODE_OK {
			return r.ReplaceVersions(ctx, credentialID, ocpiResponse.Data)
		}
	}

	return []*db.Version{}
}
