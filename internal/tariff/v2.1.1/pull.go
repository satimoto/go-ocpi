package tariff

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

func (r *TariffResolver) PullTariffsByIdentifier(ctx context.Context, countryCode string, partyID string, token string) {
	limit, offset, retries := 500, 0, 0

	if versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, "tariffs", countryCode, partyID); err == nil {
		if requestUrl, err := url.Parse(versionEndpoint.Url); err == nil {
			header := ocpi.NewOCPIRequestHeader(&token, &countryCode, &partyID)
			query := requestUrl.Query()

			if location, err := r.GetLastTariffByIdentity(ctx, countryCode, partyID); err == nil {
				query.Set("date_from", location.LastUpdated.Format(time.RFC3339Nano))
			}

			for {
				query.Set("limit", fmt.Sprintf("%d", limit))
				query.Set("offset", fmt.Sprintf("%d", offset))
				requestUrl.RawQuery = query.Encode()

				if response, err := r.OCPIRequester.Do("GET", requestUrl.String(), header, nil); err == nil {
					dto, err := r.UnmarshalPullDto(response.Body)
					limit = ocpi.GetXLimitHeader(response, limit)
					response.Body.Close()

					if err == nil && dto.StatusCode == ocpi.STATUS_CODE_OK {
						r.ReplaceTariffsByIdentifier(ctx, &countryCode, &partyID, nil, dto.Data)

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
