package version

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/rest"
	"github.com/satimoto/go-ocpi-api/util"
)

func (r *VersionDetailResolver) UpdateVersionDetail(ctx context.Context, url string, header util.OCPIRequestHeader, versionID int64) []*db.VersionEndpoint {
	if response, err := r.OCPIRequester.Do("GET", url, header, nil); err == nil {
		defer response.Close()

		ocpiResponse, err := r.UnmarshalResponse(response)

		if err == nil && ocpiResponse.StatusCode == rest.STATUS_CODE_OK {
			return r.ReplaceVersionEndpoints(ctx, versionID, ocpiResponse.Data)
		}
	}

	return []*db.VersionEndpoint{}
}
