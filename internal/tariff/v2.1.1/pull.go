package tariff

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func (r *TariffResolver) PullTariffsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string) {
	limit, offset, retries := 500, 0, 0

	if versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "tariffs", credential.CountryCode, credential.PartyID); err == nil {
		if requestUrl, err := url.Parse(versionEndpoint.Url); err == nil {
			header := transportation.NewOCPIRequestHeader(&credential.ClientToken.String, countryCode, partyID)
			query := requestUrl.Query()

			if location, err := r.GetLastTariffByIdentity(ctx, &credential.ID, countryCode, partyID); err == nil {
				query.Set("date_from", location.LastUpdated.Format(time.RFC3339Nano))
			}

			for {
				query.Set("limit", fmt.Sprintf("%d", limit))
				query.Set("offset", fmt.Sprintf("%d", offset))
				requestUrl.RawQuery = query.Encode()

				if response, err := r.OCPIRequester.Do(http.MethodGet, requestUrl.String(), header, nil); err == nil {
					dto, err := r.UnmarshalPullDto(response.Body)
					limit = transportation.GetXLimitHeader(response, limit)
					response.Body.Close()

					if err == nil && dto.StatusCode == transportation.STATUS_CODE_OK {
						r.ReplaceTariffsByIdentifier(ctx, credential, countryCode, partyID, nil, dto.Data)

						if len(dto.Data) == limit {
							offset += limit
							continue
						}
					}
				}

				retries++

				if retries >= 5 {
					break
				}
			}
		}
	}
}
