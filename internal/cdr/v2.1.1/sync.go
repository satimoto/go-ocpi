package cdr

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreCdr "github.com/satimoto/go-ocpi/internal/cdr"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *CdrResolver) SyncByIdentifier(ctx context.Context, credential db.Credential, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	log.Printf("Sync cdrs Url=%v LastUpdated=%v CountryCode=%v PartyID=%v",
		credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))
	limit, offset, retries := 500, 0, 0

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, coreCdr.IDENTIFIER, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI028", "Error retrieving version endpoint", err)
		log.Printf("OCPI028: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, coreCdr.IDENTIFIER)
		return
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI029", "Error parsing url", err)
		log.Printf("OCPI029: Url=%v", versionEndpoint.Url)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, countryCode, partyID)
	query := requestUrl.Query()

	if lastUpdated == nil {
		if cdr, err := r.GetLastCdrByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
			lastUpdated = &cdr.LastUpdated
		} else if credential.IsHub {
			oneMonthAgo := time.Now().Add(time.Hour * 24 * -30)
			lastUpdated = &oneMonthAgo
		}
	}

	if lastUpdated != nil {
		oneMonthFromLastUpdated := lastUpdated.Add(time.Hour * 24 * 30)

		if credential.IsHub && oneMonthFromLastUpdated.Before(time.Now()) {
			query.Set("date_to", oneMonthFromLastUpdated.Format(time.RFC3339))
		}

		query.Set("date_from", lastUpdated.Format(time.RFC3339))
	}

	for {
		if retries >= 5 {
			break
		}
		
		query.Set("limit", fmt.Sprintf("%d", limit))
		query.Set("offset", fmt.Sprintf("%d", offset))
		requestUrl.RawQuery = query.Encode()

		dto, limit := r.sendRequest(requestUrl.String(), header, limit)

		if dto == nil {
			retries++
			continue
		}

		retries = 0

		if dto.StatusCode == transportation.STATUS_CODE_OK {
			count := len(dto.Data)

			log.Printf("Page limit=%v offset=%v count=%v", limit, offset, count)
			r.ReplaceCdrs(ctx, credential, dto.Data)
			offset += limit

			if limit == 0 || count == 0 || count < limit {
				break
			}
		}
	}
}

func (r *CdrResolver) sendRequest(url string, header transportation.OcpiRequestHeader, limit int) (*dto.OcpiCdrsDto, int) {
	response, err := r.OcpiService.Do(http.MethodGet, url, header, nil)

	if err != nil {
		metrics.RecordError("OCPI030", "Error making request", err)
		log.Printf("OCPI030: Method=%v, Url=%v, Header=%#v", http.MethodGet, url, header)
		
		return nil, limit
	}

	dto, err := r.UnmarshalPullDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI031", "Error unmarshaling response", err)
		util.LogHttpResponse("OCPI031", url, response, true)
		
		return nil, limit
	}

	if dto.StatusCode != transportation.STATUS_CODE_OK {
		metrics.RecordError("OCPI032", "Error response failure", err)
		util.LogHttpRequest("OCPI032", url, response.Request, true)
		util.LogHttpResponse("OCPI032", url, response, true)
		log.Printf("OCPI032: StatusCode=%v, StatusMessage=%v", dto.StatusCode, dto.StatusMessage)
		
		return nil, limit
	}

	return dto, transportation.GetXLimitHeader(response, limit)
}
