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
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *TariffResolver) PullTariffsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string) {
	limit, offset, retries := 500, 0, 0

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "tariffs", credential.CountryCode, credential.PartyID)

	if err != nil {
		util.LogOnError("OCPI182", "Error retrieving version endpoint", err)
		log.Printf("OCPI182: CountryCode=%v, PartyID=%v", credential.CountryCode, credential.PartyID)
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

	if tariff, err := r.GetLastTariffByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
		query.Set("date_from", tariff.LastUpdated.Format(time.RFC3339Nano))
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
			util.LogHttpResponse("OCPI186", requestUrl.String(), response, true)
			break
		}

		retries = 0

		if dto.StatusCode == transportation.STATUS_CODE_OK {
			r.ReplaceTariffsByIdentifier(ctx, credential, countryCode, partyID, nil, dto.Data)
			offset += limit

			if len(dto.Data) < limit {
				break
			}
		}
	}
}
