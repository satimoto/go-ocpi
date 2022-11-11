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
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	coreTariff "github.com/satimoto/go-ocpi/internal/tariff"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func (r *TariffResolver) SyncByIdentifier(ctx context.Context, credential db.Credential, fullSync bool, lastUpdated *time.Time, countryCode *string, partyID *string) {
	log.Printf("Sync tariffs Url=%v LastUpdated=%v CountryCode=%v PartyID=%v",
		credential.Url, lastUpdated, util.DefaultString(countryCode, ""), util.DefaultString(partyID, ""))
	limit, offset, retries := 500, 0, 0

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, coreTariff.IDENTIFIER, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI182", "Error retrieving version endpoint", err)
		log.Printf("OCPI182: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, coreTariff.IDENTIFIER)
		return
	}

	requestUrl, err := url.Parse(versionEndpoint.Url)

	if err != nil {
		metrics.RecordError("OCPI183", "Error parsing url", err)
		log.Printf("OCPI183: Url=%v", versionEndpoint.Url)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, countryCode, partyID)
	query := requestUrl.Query()

	if !fullSync {
		if lastUpdated != nil {
			query.Set("date_from", lastUpdated.UTC().Format(time.RFC3339))
		} else if tariff, err := r.GetLastTariffByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
			query.Set("date_from", tariff.LastUpdated.Format(time.RFC3339))
		}
	}

	for {
		query.Set("limit", fmt.Sprintf("%d", limit))
		query.Set("offset", fmt.Sprintf("%d", offset))
		requestUrl.RawQuery = query.Encode()

		response, err := r.OcpiService.Do(http.MethodGet, requestUrl.String(), header, nil)

		if err != nil {
			metrics.RecordError("OCPI184", "Error making request", err)
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
			metrics.RecordError("OCPI185", "Error unmarshaling response", err)
			util.LogHttpResponse("OCPI185", requestUrl.String(), response, true)
			break
		}

		limit = transportation.GetXLimitHeader(response, limit)

		if dto.StatusCode != transportation.STATUS_CODE_OK {
			metrics.RecordError("OCPI186", "Error response failure", err)
			util.LogHttpRequest("OCPI186", requestUrl.String(), response.Request, true)
			util.LogHttpResponse("OCPI186", requestUrl.String(), response, true)
			log.Printf("OCPI186: StatusCode=%v, StatusMessage=%v", dto.StatusCode, dto.StatusMessage)
			break
		}

		retries = 0

		if dto.StatusCode == transportation.STATUS_CODE_OK {
			count := len(dto.Data)

			log.Printf("Page limit=%v offset=%v count=%v", limit, offset, count)
			r.ReplaceTariffsByIdentifier(ctx, credential, countryCode, partyID, nil, dto.Data)
			offset += limit

			if limit == 0 || count == 0 || count < limit {
				break
			}
		}
	}
}
