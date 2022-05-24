package version

import (
	"context"
	"log"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *VersionResolver) ReplaceVersions(ctx context.Context, credentialID int64, dto []*VersionDto) []*db.Version {
	versions := []*db.Version{}

	if dto != nil {
		r.Repository.DeleteVersions(ctx, credentialID)

		for _, versionDto := range dto {
			versionParams := NewCreateVersionParams(credentialID, versionDto)
			version, err := r.Repository.CreateVersion(ctx, versionParams)

			if err != nil {
				util.LogOnError("OCPI211", "Error creating version", err)
				log.Printf("OCPI211: Params=%#v", versionParams)
				continue
			}

			versions = append(versions, &version)
		}
	}

	return versions
}

func (r *VersionResolver) PullVersions(ctx context.Context, url string, header transportation.OcpiRequestHeader, credentialID int64) []*db.Version {
	response, err := r.OcpiRequester.Do(http.MethodGet, url, header, nil)

	if err != nil {
		util.LogOnError("OCPI212", "Error making request", err)
		log.Printf("OCPI212: Method=%v, Url=%v, Header=%#v", http.MethodGet, url, header)
		return []*db.Version{}
	}

	ocpiResponse, err := r.UnmarshalPullDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		util.LogOnError("OCPI213", "Error unmarshalling response", err)
		util.LogHttpResponse("OCPI213", url, response, true)
		return []*db.Version{}
	}

	if ocpiResponse.StatusCode != transportation.STATUS_CODE_OK {
		util.LogOnError("OCPI214", "Error response failure", err)
		util.LogHttpResponse("OCPI214", url, response, true)
		return []*db.Version{}
	}

	return r.ReplaceVersions(ctx, credentialID, ocpiResponse.Data)
}
