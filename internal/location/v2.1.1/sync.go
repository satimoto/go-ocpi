package location

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *LocationResolver) SyncByIdentifier(ctx context.Context, credential db.Credential, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	log.Printf("Sync locations Url=%v LastUpdated=%v CountryCode=%v PartyID=%v",
		credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))
	limit, offset, retries := 500, 0, 0
	identifier := "locations"

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI125", "Error retrieving version endpoint", err)
		log.Printf("OCPI125: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
		return
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI126", "Error parsing url", err)
		log.Printf("OCPI126: Url=%v", versionEndpoint.Url)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, countryCode, partyID)
	query := requestUrl.Query()

	if !fullSync {
		if lastUpdated != nil {
			query.Set("date_from", lastUpdated.UTC().Format(time.RFC3339))
		} else if location, err := r.GetLastLocationByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
			query.Set("date_from", location.LastUpdated.Format(time.RFC3339))
		}
	}

	for {
		query.Set("limit", fmt.Sprintf("%d", limit))
		query.Set("offset", fmt.Sprintf("%d", offset))
		requestUrl.RawQuery = query.Encode()

		response, err := r.OcpiService.Do(http.MethodGet, requestUrl.String(), header, nil)

		if err != nil {
			metrics.RecordError("OCPI127", "Error making request", err)
			log.Printf("OCPI127: Method=%v, Url=%v, Header=%#v", http.MethodGet, requestUrl.String(), header)
			retries++

			if retries >= 5 {
				break
			}

			continue
		}

		dto, err := r.UnmarshalPullDto(response.Body)
		defer response.Body.Close()

		if err != nil {
			metrics.RecordError("OCPI128", "Error unmarshalling response", err)
			util.LogHttpResponse("OCPI128", requestUrl.String(), response, true)
			break
		}

		limit = transportation.GetXLimitHeader(response, limit)

		if dto.StatusCode != transportation.STATUS_CODE_OK {
			metrics.RecordError("OCPI129", "Error response failure", err)
			util.LogHttpRequest("OCPI129", versionEndpoint.Url, response.Request, true)
			util.LogHttpResponse("OCPI129", requestUrl.String(), response, true)
			log.Printf("OCPI129: StatusCode=%v, StatusMessage=%v", dto.StatusCode, dto.StatusMessage)
			break
		}

		retries = 0

		if dto.StatusCode == transportation.STATUS_CODE_OK {
			r.ReplaceLocations(ctx, credential, dto.Data)
			offset += limit

			if limit == 0 || len(dto.Data) < limit {
				break
			}
		}
	}
}
