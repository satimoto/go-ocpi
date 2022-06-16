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
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *CdrResolver) PullCdrsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string) {
	limit, offset, retries := 500, 0, 0
	identifier := "cdrs"

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx,identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		util.LogOnError("OCPI028", "Error retrieving version endpoint", err)
		log.Printf("OCPI028: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
		return
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		util.LogOnError("OCPI029", "Error parsing url", err)
		log.Printf("OCPI029: Url=%v", versionEndpoint.Url)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, countryCode, partyID)
	query := requestUrl.Query()

	if cdr, err := r.GetLastCdrByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
		query.Set("date_from", cdr.LastUpdated.Format(time.RFC3339Nano))
	}

	for {
		query.Set("limit", fmt.Sprintf("%d", limit))
		query.Set("offset", fmt.Sprintf("%d", offset))
		requestUrl.RawQuery = query.Encode()

		response, err := r.OcpiRequester.Do(http.MethodGet, requestUrl.String(), header, nil)

		if err != nil {
			util.LogOnError("OCPI030", "Error making request", err)
			log.Printf("OCPI030: Method=%v, Url=%v, Header=%#v", http.MethodGet, requestUrl.String(), header)
			retries++

			if retries >= 5 {
				break
			}

			continue
		}

		dto, err := r.UnmarshalPullDto(response.Body)
		defer response.Body.Close()

		if err != nil {
			util.LogOnError("OCPI031", "Error unmarshalling response", err)
			util.LogHttpResponse("OCPI031", requestUrl.String(), response, true)
			break
		}

		limit = transportation.GetXLimitHeader(response, limit)

		if dto.StatusCode != transportation.STATUS_CODE_OK {
			util.LogOnError("OCPI032", "Error response failure", err)
			util.LogHttpRequest("OCPI032", requestUrl.String(), response.Request, true)
			util.LogHttpResponse("OCPI032", requestUrl.String(), response, true)
			break
		}

		retries = 0
	
		if dto.StatusCode == transportation.STATUS_CODE_OK {
			r.ReplaceCdrsByIdentifier(ctx, credential, countryCode, partyID, dto.Data)
			offset += limit

			if len(dto.Data) < limit {
				break
			}
		}
	}
}
