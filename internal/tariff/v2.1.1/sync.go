package tariff

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TariffResolver) SyncByIdentifier(ctx context.Context, credential db.Credential, lastUpdated *time.Time, countryCode *string, partyID *string) {
	log.Printf("Sync tariffs Url=%v LastUpdated=%v CountryCode=%v PartyID=%v", credential.Url, lastUpdated, countryCode, partyID)
	limit, offset, retries := 500, 0, 0
	identifier := "tariffs"

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, identifier, credential.CountryCode, credential.PartyID)

	if err != nil {
		util.LogOnError("OCPI182", "Error retrieving version endpoint", err)
		log.Printf("OCPI182: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, identifier)
		return
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		util.LogOnError("OCPI183", "Error parsing url", err)
		log.Printf("OCPI183: Url=%v", versionEndpoint.Url)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, countryCode, partyID)
	query := requestUrl.Query()

	if lastUpdated != nil {
		query.Set("date_from", lastUpdated.UTC().Format(time.RFC3339))
	} else if tariff, err := r.GetLastTariffByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
		query.Set("date_from", tariff.LastUpdated.Format(time.RFC3339))
	}

	for {
		query.Set("limit", fmt.Sprintf("%d", limit))
		query.Set("offset", fmt.Sprintf("%d", offset))
		requestUrl.RawQuery = query.Encode()

		response, err := r.OcpiRequester.Do(http.MethodGet, requestUrl.String(), header, nil)

		if err != nil {
			util.LogOnError("OCPI184", "Error making request", err)
			log.Printf("OCPI184: Method=%v, Url=%v, Header=%#v", http.MethodGet, requestUrl.String(), header)
			retries++

			if retries >= 5 {
				break
			}

			continue
		}

		dto, err := r.UnmarshalPullDto(response.Body)
		response.Body.Close()

		if err != nil {
			util.LogOnError("OCPI185", "Error unmarshalling response", err)
			util.LogHttpResponse("OCPI185", requestUrl.String(), response, true)
			break
		}

		limit = transportation.GetXLimitHeader(response, limit)

		if dto.StatusCode != transportation.STATUS_CODE_OK {
			util.LogOnError("OCPI186", "Error response failure", err)
			util.LogHttpRequest("OCPI186", requestUrl.String(), response.Request, true)
			util.LogHttpResponse("OCPI186", requestUrl.String(), response, true)
			log.Printf("OCPI186: StatusCode=%v, StatusMessage=%v", dto.StatusCode, dto.StatusMessage)
			break
		}

		retries = 0

		if dto.StatusCode == transportation.STATUS_CODE_OK {
			r.ReplaceTariffsByIdentifier(ctx, credential, countryCode, partyID, nil, dto.Data)
			offset += limit

			if limit == 0 || len(dto.Data) < limit {
				break
			}
		}
	}
}
