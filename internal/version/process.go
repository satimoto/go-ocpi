package version

import (
	"context"
	"log"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *VersionResolver) ReplaceVersions(ctx context.Context, credentialID int64, versionsDto []*coreDto.VersionDto) []*db.Version {
	versions := []*db.Version{}

	if versionsDto != nil {
		r.Repository.DeleteVersions(ctx, credentialID)

		for _, versionDto := range versionsDto {
			versionParams := NewCreateVersionParams(credentialID, versionDto)
			version, err := r.Repository.CreateVersion(ctx, versionParams)

			if err != nil {
				metrics.RecordError("OCPI211", "Error creating version", err)
				log.Printf("OCPI211: Params=%#v", versionParams)
				continue
			}

			versions = append(versions, &version)
		}
	}

	return versions
}

func (r *VersionResolver) PullVersions(ctx context.Context, url string, header transportation.OcpiRequestHeader, credentialID int64) []*db.Version {
	response, err := r.OcpiService.Do(http.MethodGet, url, header, nil)

	if err != nil {
		metrics.RecordError("OCPI212", "Error making request", err)
		log.Printf("OCPI212: Method=%v, Url=%v, Header=%#v", http.MethodGet, url, header)
		return []*db.Version{}
	}

	pullDto, err := r.UnmarshalPullDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI213", "Error unmarshaling response", err)
		util.LogHttpResponse("OCPI213", url, response, true)
		return []*db.Version{}
	}

	if pullDto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI214", "Error response failure", err)
		util.LogHttpRequest("OCPI214", url, response.Request, true)
		util.LogHttpResponse("OCPI214", url, response, true)
		log.Printf("OCPI214: StatusCode=%v, StatusMessage=%v", pullDto.StatusCode, pullDto.StatusMessage)
		return []*db.Version{}
	}

	return r.ReplaceVersions(ctx, credentialID, pullDto.Data)
}
